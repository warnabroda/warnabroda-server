package routes

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/mostafah/mandrill"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetWarnings(enc Encoder, db gorp.SqlExecutor) (int, string) {
	var warnings []models.Warning
	_, err := db.Select(&warnings, "select * from warnings order by id")
	if err != nil {
		checkErr(err, "select failed")
		return http.StatusInternalServerError, ""
	}
	return http.StatusOK, Must(enc.Encode(warningsToIface(warnings)...))
}

func GetWarning(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	obj, _ := db.Get(models.Warning{}, id)
	if err != nil || obj == nil {
		checkErr(err, "get failed")
		// Invalid id, or does not exist
		return http.StatusNotFound, ""
	}
	entity := obj.(*models.Warning)
	return http.StatusOK, Must(enc.EncodeOne(entity))
}

func AddWarning(entity models.Warning, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {

	entity.Sent = false
	entity.Created_by = "system"
	entity.Created_date = models.JDate(time.Now())
	entity.Lang_key = "br"

	err := db.Insert(&entity)
	if err != nil {
		checkErr(err, "insert failed")
		return http.StatusConflict, ""
	}
	w.Header().Set("Location", fmt.Sprintf("/warnabroda/warnings/%d", entity.Id))
	if entity.Id_contact_type == 1 {
		go sendEmail(&entity, db)
	}
	return http.StatusCreated, Must(enc.EncodeOne(entity))
}

func sendEmail(entity *models.Warning, db gorp.SqlExecutor) {
	fmt.Println("Mailling Routine Begin, 世界")

	mandrill.Key = "qUX983QXREtaLojEpJyxmw"
	// you can test your API key with Ping
	err := mandrill.Ping()
	// everything is OK if err is nil

	template_byte, err := ioutil.ReadFile("./models/warning.html")
	checkErr(err, "File Opening ERROR")
	template_email_string := string(template_byte[:])
	fmt.Println("OLA")

	subject := GetRandomSubject(db)

	message := SelectMessage(db, entity.Id_message)
	fmt.Println("CHEGOU 1")

	var email_content string
	email_content = strings.Replace(template_email_string, "{{warning}}", message.Name, 1)
	fmt.Println("CHEGOU 2")

	msg := mandrill.NewMessageTo(entity.Contact, subject.Name)
	fmt.Println("CHEGOU 3")
	msg.HTML = email_content
	// msg.Text = "plain text content" // optional
	msg.Subject = subject.Name
	msg.FromEmail = "warnabroda@gmail.com"
	msg.FromName = "Warn A Broda: Dá um toque"

	//envio assincrono = true // envio sincrono = false
	res, err := msg.Send(false)

	if res[0] != nil {
		UpdateWarningSent(entity, db)
	} else {
		fmt.Println(res[0])
	}
	fmt.Println("Mailling Routine End, 世界")
}

func UpdateWarningSent(entity *models.Warning, db gorp.SqlExecutor) {
	entity.Sent = true
	entity.Last_modified_date = models.JDate(time.Now())
	_, err := db.Update(entity)
	if err != nil {
		checkErr(err, "update failed")
	}
}

func UpdateWarning(entity models.Warning, enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	obj, _ := db.Get(models.Warning{}, id)
	if err != nil || obj == nil {
		checkErr(err, "get failed")
		// Invalid id, or does not exist
		return http.StatusNotFound, ""
	}
	oldEntity := obj.(*models.Warning)

	entity.Id = oldEntity.Id
	_, err = db.Update(&entity)
	if err != nil {
		checkErr(err, "update failed")
		return http.StatusConflict, ""
	}
	return http.StatusOK, Must(enc.EncodeOne(entity))
}

func DeleteWarning(db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	obj, _ := db.Get(models.Warning{}, id)
	if err != nil || obj == nil {
		checkErr(err, "get failed")
		// Invalid id, or does not exist
		return http.StatusNotFound, ""
	}
	entity := obj.(*models.Warning)
	_, err = db.Delete(entity)
	if err != nil {
		checkErr(err, "delete failed")
		return http.StatusConflict, ""
	}
	return http.StatusNoContent, ""
}

func warningsToIface(v []models.Warning) []interface{} {
	if len(v) == 0 {
		return nil
	}
	ifs := make([]interface{}, len(v))
	for i, v := range v {
		ifs[i] = v
	}
	return ifs
}
