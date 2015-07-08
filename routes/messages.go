package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessionauth"
	"gitlab.com/warnabroda/warnabrodagomartini/models"
)

const (
	SQL_MESSAGES_BY_LANG_KEY = "SELECT id, name, lang_key, active FROM messages WHERE lang_key=? AND active=true ORDER BY name"
	SQL_MESSAGES_BY_ID       = "SELECT id, name, lang_key, active FROM messages WHERE id=?"
	SQL_MESSAGES_ALL         = "SELECT " +
		" 	DISTINCT(m.name) as name, " +
		"	m.id as id, " +
		"	m.lang_key as lang_key, 	" +
		"	m.active as active, 	" +
		"	COUNT(w.id) as total, " +
		"	(SELECT COUNT(*) FROM warnings ww WHERE ww.sent = false AND ww.Id_message = m.id) AS not_sent, " +
		"	(SELECT COUNT(*) FROM warnings www WHERE www.sent = true AND www.Id_message = m.id) AS sent, " +
		"	(SELECT COUNT(*) FROM warnings wwww WHERE wwww.Id_contact_type = 1 AND wwww.Id_message = m.id) AS email, " +
		"	(SELECT COUNT(*) FROM warnings wwwww WHERE wwwww.Id_contact_type = 2 AND wwwww.Id_message = m.id) AS sms, " +
		"	(SELECT COUNT(*) FROM warnings wwwwww WHERE wwwwww.Id_contact_type = 3 AND wwwwww.Id_message = m.id) AS whatsapp " +
		" FROM messages m " +
		" LEFT JOIN warnings w on w.id_message = m.id " +
		" GROUP BY m.name " +
		" ORDER BY total DESC, m.Lang_key DESC, m.name ASC "
	SQL_RANDOM_MESSAGES = " SELECT m.id, m.name, m.lang_key, m.active FROM warnabroda.messages AS m JOIN  " +
		" (SELECT CEIL(RAND() * (SELECT MAX(id) FROM warnabroda.messages)) AS id) AS m2 " +
		" WHERE m.lang_key=? AND m.active=true AND m.id >= m2.id " +
		" ORDER BY m.id ASC LIMIT 7 "
)

var messagesEn []models.DefaultStruct
var messagesEs []models.DefaultStruct
var messagesPt []models.DefaultStruct

func init() {
	_, errEn := models.Dbm.Select(&messagesEn, SQL_MESSAGES_BY_LANG_KEY, "en")
	checkErr(errEn, "select failed")

	_, errEs := models.Dbm.Select(&messagesEs, SQL_MESSAGES_BY_LANG_KEY, "es")
	checkErr(errEs, "select failed")

	_, errPt := models.Dbm.Select(&messagesPt, SQL_MESSAGES_BY_LANG_KEY, "pt-br")
	checkErr(errPt, "select failed")

}

func GetMessages(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	var messages []models.DefaultStruct
	lang_key := parms["lang_key"]

	_, err := db.Select(&messages, SQL_MESSAGES_BY_LANG_KEY, lang_key)
	if err != nil {
		checkErr(err, "select failed")
		return http.StatusBadRequest, ""
	}

	return http.StatusOK, Must(enc.Encode(messagesToIface(messages)...))
}

func GetMessagesStats(enc Encoder, db gorp.SqlExecutor, user sessionauth.User) (int, string) {

	u := UserById(user.UniqueId().(int), db)

	if user.IsAuthenticated() && u.UserRole == models.ROLE_ADMIN {
		var messages []models.Messages

		_, err := db.Select(&messages, SQL_MESSAGES_ALL)
		if err != nil {
			checkErr(err, "select failed")
			return http.StatusBadRequest, ""
		}

		return http.StatusOK, Must(enc.Encode(messagesToIfaceM(messages)...))
	}

	return http.StatusUnauthorized, ""
}

func GetMessage(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])

	if err != nil {
		// Invalid id, or does not exist
		return http.StatusBadRequest, ""
	}

	obj := models.MessageStruct{}
	err = db.SelectOne(&obj, "SELECT * FROM messages WHERE id=?", id)

	if err != nil {
		checkErr(err, "get failed")
		// Invalid id, or does not exist
		return http.StatusBadRequest, ""
	}
	return http.StatusOK, Must(enc.EncodeOne(obj))
}

func SelectMessage(db gorp.SqlExecutor, id int64) models.DefaultStruct {

	entity := models.DefaultStruct{}

	err := db.SelectOne(&entity, SQL_MESSAGES_BY_ID, id)

	if err != nil {
		checkErr(err, "select failed")
	}

	return entity
}

func AddMessage(entity models.DefaultStruct, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {
	err := db.Insert(&entity)
	if err != nil {
		checkErr(err, "insert failed")
		return http.StatusBadRequest, ""
	}
	w.Header().Set("Location", fmt.Sprintf("/warnabroda/messages/%d", entity.Id))
	return http.StatusCreated, Must(enc.EncodeOne(entity))
}

func SaveOrUpdateMessage(entity models.MessageStruct, enc Encoder, db gorp.SqlExecutor, user sessionauth.User) (int, string) {

	u := UserById(user.UniqueId().(int), db)

	if user.IsAuthenticated() && u.UserRole == models.ROLE_ADMIN {

		entity.Last_modified_by = user.UniqueId().(int)

		if entity.Id < 1 {
			err := db.Insert(&entity)
			if err != nil {
				checkErr(err, "insert failed")
				return http.StatusBadRequest, ""
			}
		} else {
			obj, _ := db.Get(models.MessageStruct{}, entity.Id)
			if obj == nil {
				// Invalid id, or does not exist
				return http.StatusBadRequest, ""
			}

			_, err := db.Update(&entity)
			if err != nil {
				checkErr(err, "update failed")
				return http.StatusBadRequest, ""
			}
		}

		return http.StatusOK, Must(enc.EncodeOne(entity))

	}

	return http.StatusUnauthorized, ""
}

func DeleteMessage(db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	obj, _ := db.Get(models.DefaultStruct{}, id)
	if err != nil || obj == nil {
		checkErr(err, "get failed")
		// Invalid id, or does not exist
		return http.StatusBadRequest, ""
	}
	entity := obj.(*models.DefaultStruct)
	_, err = db.Delete(entity)
	if err != nil {
		checkErr(err, "delete failed")
		return http.StatusBadRequest, ""
	}
	return http.StatusOK, ""
}

func messagesToIface(v []models.DefaultStruct) []interface{} {
	if len(v) == 0 {
		return nil
	}
	ifs := make([]interface{}, len(v))
	for i, v := range v {
		ifs[i] = v
	}
	return ifs
}

func messagesToIfaceM(v []models.Messages) []interface{} {
	if len(v) == 0 {
		return nil
	}
	ifs := make([]interface{}, len(v))
	for i, v := range v {
		ifs[i] = v
	}
	return ifs
}

func GetRandomMessagesByLanguage(amount int, lang_key string, db gorp.SqlExecutor) []models.DefaultStruct {
	fmt.Println("GetRandomMessagesByLanguage: " + strconv.Itoa(amount) + ", " + lang_key)

	messages := make([]models.DefaultStruct, amount)

	index := 0
	for index != amount {
		if lang_key == "pt-br" {
			messages[index] = messagesPt[index]
		}
		if lang_key == "en" {
			messages[index] = messagesEn[index]
		}
		if lang_key == "es" {
			messages[index] = messagesEs[index]
		}
		index++

	}

	return messages
}
