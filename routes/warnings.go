package routes

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	//	"io/ioutil"
	//	"os"

	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessionauth"
	"github.com/warnabroda/warnabroda-server/messages"
	"github.com/warnabroda/warnabroda-server/models"
)

const (
	SQL_WARNING_BYID              = "SELECT * FROM warnings ORDER BY id"
	SQL_SELECT_REPLY_BY_READ_HASH = "SELECT * FROM warning_resp WHERE read_hash = :hash"
	SQL_SELECT_REPLY_BY_RESP_HASH = "SELECT * FROM warning_resp WHERE resp_hash = :hash"
	SQL_CHECK_SENT_WARN           = "SELECT COUNT(*) FROM warnings " +
		" WHERE Id_contact_type = :id_contact_type AND Sent = true AND " +
		" (Contact = :contact OR Ip LIKE :ip ) AND " +
		" Created_date BETWEEN :lower_str_date AND :upper_str_date AND " +
		" Id <> :id "
)

func BuildCountWarningsSql(count_by string) string {
	fmt.Println("BuildCountWarningsSql")

	sql := " SELECT COUNT(*) FROM warnabroda.warnings "
	sql += " WHERE sent = :sent AND (created_date + INTERVAL :interval HOUR) > NOW()"

	switch count_by {
	case "ip":
		sql += " AND id_contact_type = :id_contact_type "
		sql += " AND ip = :ip "
	case "same_message_by_ip":
		sql += " AND contact = :contact "
		sql += " AND id_message = :id_message "
		sql += " AND ip = :ip "
	case "same_message":
		sql += " AND contact = :contact "
		sql += " AND id_message = :id_message "
		sql += " AND ip <> :ip "
	}

	return sql

}

func GetWarnings(enc Encoder, db gorp.SqlExecutor) (int, string) {

	var warnings []models.Warning
	_, err := db.Select(&warnings, SQL_WARNING_BYID)
	checkErr(err, "LIST ALL WARNINGS ERROR")

	if err != nil {
		return http.StatusInternalServerError, ""
	}
	return http.StatusOK, Must(enc.Encode(warningsToIface(warnings)...))
}

func GetWarning(id int64, db gorp.SqlExecutor) *models.Warning {

	obj, err := db.Get(models.Warning{}, id)

	if err != nil || obj == nil {
		return nil
	}

	entity := obj.(*models.Warning)
	return entity
}

func GetWarningDetail(enc Encoder, db gorp.SqlExecutor, user sessionauth.User, parms martini.Params) (int, string) {

	u := UserById(user.UniqueId().(int), db)

	if user.IsAuthenticated() && u.UserRole == models.ROLE_ADMIN {

		id, err := strconv.Atoi(parms["id"])
		obj, _ := db.Get(models.Warning{}, id)
		if err != nil || obj == nil {
			checkErr(err, "GET WARNING DETAIL FAILED")
			// Invalid id, or does not exist
			return http.StatusNotFound, ""
		}
		entity := obj.(*models.Warning)
		return http.StatusOK, Must(enc.EncodeOne(entity))
	}

	return http.StatusUnauthorized, ""

}

func UpdateWarningSent(entity *models.Warning, db gorp.SqlExecutor) bool {
	entity.Sent = true
	entity.Last_modified_date = time.Now().String()
	_, err := db.Update(entity)
	checkErr(err, "ERROR UpdateWarningSent ERROR")
	return err == nil
}

func ConfirmWarning(entity models.DefaultStruct, enc Encoder, db gorp.SqlExecutor) (int, string) {

	status := &models.DefaultStruct{
		Id:       http.StatusNotFound,
		Name:     "Warning Update Failed.",
		Lang_key: "en",
	}

	warning := GetWarning(entity.Id, db)

	if warning != nil {

		warning.Message = entity.Name
		if UpdateWarningSent(warning, db) {
			status.Id = http.StatusAccepted
			status.Name = "Warning Update Success."
			status.Lang_key = warning.Lang_key
		}

	}

	return http.StatusOK, Must(enc.EncodeOne(status))
}

func isInvalidWarning(entity *models.Warning) bool {
	fmt.Println("isInvalidWarning")
	return (entity.Id_message < 1) || (entity.Id_contact_type < 1) || (len(entity.Contact) < 5) || (len(entity.Lang_key) < 2) || (len(entity.Ip) < 6)

}

// Receives a warning tru, inserts the request and process the warning and then respond to the interface
//TODO: use (session sessions.Session, r *http.Request) to prevent flood
func AddWarning(entity models.Warning, enc Encoder, db gorp.SqlExecutor, r *http.Request) (int, string) {
	fmt.Println("AddWarning")
	if isInvalidWarning(&entity) {
		return http.StatusBadRequest, Must(enc.EncodeOne(entity))
	}

	status := &models.DefaultStruct{
		Id:       http.StatusOK,
		Name:     messages.GetLocaleMessage(entity.Lang_key, "MSG_WARNING_SENT_SUCCESS"),
		Lang_key: entity.Lang_key,
		Type:     models.MSG_TYPE_WARNING,
	}

	entity.Sent = false
	entity.Created_by = "system"
	//entity.Created_date = time.Now().String()

	err := db.Insert(&entity)
	checkErr(err, "INSERT WARNING ERROR")
	if err != nil {
		return http.StatusBadRequest, Must(enc.EncodeOne(entity))
	}

	ingnored := InIgnoreList(db, entity.Contact)

	if ingnored != nil && ingnored.Confirmed {
		status = &models.DefaultStruct{
			Id:       http.StatusBadRequest,
			Name:     messages.GetLocaleMessage(entity.Lang_key, "MSG_IGNORED_USER"),
			Lang_key: entity.Lang_key,
			Type:     models.MSG_TYPE_WARNING,
		}
	} else {
		processWarn(&entity, db, status)
	}

	return http.StatusCreated, Must(enc.EncodeOne(status))
}

// After registered in the Database, the warn is processed in order to verify:
// - @isSameWarnSentByIp
// - @isSameWarnSentTwiceOrMoreDifferentIp
// - if none of above occurs the warn is processed by its type(Email, SMS, Whatsapp, etc...)
//		- @routers.email.ProcessEmail
//		- @routers.sms.ProcessSMS
func processWarn(warning *models.Warning, db gorp.SqlExecutor, status *models.DefaultStruct) {
	fmt.Println("processWarn")

	status.Lang_key = warning.Lang_key
	if isSameWarnSentByIp(warning, db) {
		status.Id = http.StatusBadRequest
		status.Name = strings.Replace(messages.GetLocaleMessage(warning.Lang_key, "MSG_SMS_SAME_WARN_BY_IP"), "{{ip}}", warning.Ip, 1)
		status.Name = strings.Replace(status.Name, "{{time}}", "2", 1)
	} else if isSameWarnSentTwiceOrMoreDifferentIp(warning, db) {
		status.Id = http.StatusBadRequest
		status.Name = strings.Replace(messages.GetLocaleMessage(warning.Lang_key, "MSG_SMS_SAME_WARN_DIFF_IP"), "{{time}}", "2", 1)
	} else {
		if warning.WarnResp != nil && warning.WarnResp.Reply_to != "" {
			ProcessWarnReply(warning, db)
		} else {
			warning.WarnResp = nil
		}

		switch warning.Id_contact_type {
		case 1:
			go ProcessEmail(warning, db)
		case 2:
			ProcessSMS(warning, db, status)
		case 3:
			go ProcessWhatsapp(warning, db)
		default:
			return
		}

	}
}

func ProcessWarnReply(warning *models.Warning, db gorp.SqlExecutor) {
	fmt.Println("ProcessWarnReply")

	warning.WarnResp.Id_warning = warning.Id

	warning.WarnResp.Lang_key = warning.Lang_key
	warning.WarnResp.Resp_hash = GenerateSha1(warning.Contact + "-" + warning.Created_date)
	warning.WarnResp.Read_hash = GenerateSha1(warning.WarnResp.Reply_to + "-" + warning.Created_date)
	warning.WarnResp.Reply_to = warning.WarnResp.Reply_to
	warning.WarnResp.Created_date = warning.Created_date

	if strings.Contains(warning.WarnResp.Reply_to, "@") {
		warning.WarnResp.Id_contact_type = 1
	} else {
		warning.WarnResp.Id_contact_type = 3
	}

	err := db.Insert(warning.WarnResp)
	checkErr(err, "INSERT WARNING ERROR")

}

func GenerateSha1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))

	byteStr := hash.Sum(nil)

	return fmt.Sprintf("%x", byteStr)

}

// return true if a warn, with same message and same ip, attempts to be sent, if so respond back to interface denying the service;
func isSameWarnSentByIp(warning *models.Warning, db gorp.SqlExecutor) bool {
	fmt.Println("isSameWarnSentByIp")

	exists, err := db.SelectInt(BuildCountWarningsSql("same_message_by_ip"), map[string]interface{}{
		"sent":       true,
		"contact":    warning.Contact,
		"interval":   2,
		"id_message": warning.Id_message,
		"ip":         warning.Ip,
	})
	checkErr(err, "SELECT isSameWarnSentByIp ERROR")

	return exists >= 1
}

// return true if a warn, with same message and different ip, attempts to be sent more than twice, if so respond back to interface denying the service;
func isSameWarnSentTwiceOrMoreDifferentIp(warning *models.Warning, db gorp.SqlExecutor) bool {
	fmt.Println("isSameWarnSentTwiceOrMoreDifferentIp")

	exists, err := db.SelectInt(BuildCountWarningsSql("same_message"), map[string]interface{}{
		"sent":       true,
		"contact":    warning.Contact,
		"interval":   2,
		"id_message": warning.Id_message,
		"ip":         warning.Ip,
	})
	checkErr(err, "SELECT isSameWarnSentTwiceOrMoreDifferentIp ERROR")

	return exists >= 2
}

//turns the warning struct into an interface
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

func GetReplyById(id int64, db gorp.SqlExecutor) *models.WarningResp {

	obj, err := db.Get(models.WarningResp{}, id)

	if err != nil || obj == nil {
		return nil
	}

	entity := obj.(*models.WarningResp)
	return entity
}

func UpdateReplySent(entity *models.WarningResp, db gorp.SqlExecutor) bool {
	entity.Sent = true

	_, err := db.Update(entity)
	checkErr(err, "ERROR UpdateReplySent ERROR")
	return err == nil
}

func GetReplyByHash(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {

	var warning *models.Warning
	hash := parms["hash"]

	respReply := getReplyRespHash(hash, db)
	readReply := getReplyReadHash(hash, db)

	if respReply == nil && readReply == nil {
		fmt.Println("FAILED: neither resp or read hash matches")
		return http.StatusBadRequest, ""

	} else if respReply != nil {

		respReply.Read_hash = ""
		warning = GetWarning(respReply.Id_warning, db)
		warning.WarnResp = respReply

	} else if readReply != nil {

		readReply.Resp_hash = ""
		warning = GetWarning(readReply.Id_warning, db)
		warning.WarnResp = readReply

	}

	clearReturn(warning)

	return http.StatusOK, Must(enc.EncodeOne(warning))

}

func getReplyRespHash(hash string, db gorp.SqlExecutor) *models.WarningResp {

	var reply models.WarningResp
	err := db.SelectOne(&reply, SQL_SELECT_REPLY_BY_RESP_HASH,
		map[string]interface{}{
			"hash": hash,
		})
	if err != nil {
		return nil
	}

	return &reply
}

func getReplyReadHash(hash string, db gorp.SqlExecutor) *models.WarningResp {
	var reply models.WarningResp
	err := db.SelectOne(&reply, SQL_SELECT_REPLY_BY_READ_HASH,
		map[string]interface{}{
			"hash": hash,
		})
	if err != nil {
		return nil
	}

	return &reply
}

func clearReturn(entity *models.Warning) {
	entity.Message = ""
	entity.Ip = ""
	entity.Browser = ""
	entity.Operating_system = ""
	entity.Device = ""
	entity.Raw = ""
	entity.Created_by = ""
	entity.Last_modified_by = ""
	entity.Last_modified_date = ""

	entity.WarnResp.Reply_to = ""
	entity.WarnResp.Ip = ""
	entity.WarnResp.Browser = ""
	entity.WarnResp.Operating_system = ""
	entity.WarnResp.Device = ""
	entity.WarnResp.Raw = ""
	if entity.WarnResp.Reply_date == "0000-00-00 00:00:00" {
		entity.WarnResp.Reply_date = ""
	}

	if entity.WarnResp.Response_read == "0000-00-00 00:00:00" {
		entity.WarnResp.Response_read = ""
	}
}

func isInvalidReply(entity *models.WarningResp) bool {
	return (entity.Id < 1) || (len(entity.Resp_hash) < 10) || (len(entity.Ip) < 6)
}

func SetReply(entity models.WarningResp, enc Encoder, db gorp.SqlExecutor) (int, string) {

	if isInvalidReply(&entity) {
		return http.StatusBadRequest, Must(enc.EncodeOne(entity))
	}

	obj, err := db.Get(models.WarningResp{}, entity.Id)
	replyObj := obj.(*models.WarningResp)

	if err != nil || replyObj == nil || entity.Resp_hash != replyObj.Resp_hash {
		return http.StatusBadRequest, ""
	} else {
		replyObj.Message = entity.Message
		replyObj.Ip = entity.Ip
		replyObj.Browser = entity.Browser
		replyObj.Operating_system = entity.Operating_system
		replyObj.Device = entity.Device
		replyObj.Raw = entity.Raw
		replyObj.Reply_date = entity.Reply_date
		replyObj.Timezone = entity.Timezone

		go notifyReplyDone(replyObj, db)

		_, err = db.Update(replyObj)
		checkErr(err, "ERROR UpdateWarningSent ERROR")
	}

	return http.StatusOK, Must(enc.EncodeOne(replyObj))
}

func notifyReplyDone(entity *models.WarningResp, db gorp.SqlExecutor) {
	obj, err := db.Get(models.Warning{}, entity.Id_warning)
	if err == nil {
		warningObj := obj.(*models.Warning)
		warningObj.WarnResp = entity

		if strings.Contains(entity.Reply_to, "@") { //notify the reply is ready via e-mail
			SendEmailReplyDone(warningObj, db)
		} else { //notify the reply is ready via e-mail
			SendWhatsappReplyDone(warningObj, db)
		}
	}

}

func isInvalidReplyRead(entity *models.WarningResp) bool {
	return (entity.Id < 1) || (len(entity.Read_hash) < 10)
}

func ReadReply(entity models.WarningResp, enc Encoder, db gorp.SqlExecutor) (int, string) {

	if isInvalidReplyRead(&entity) {
		return http.StatusBadRequest, Must(enc.EncodeOne(entity))
	}

	obj, err := db.Get(models.WarningResp{}, entity.Id)
	replyObj := obj.(*models.WarningResp)

	if err != nil || replyObj == nil || entity.Read_hash != replyObj.Read_hash {
		return http.StatusBadRequest, ""
	} else {
		replyObj.Response_read = entity.Response_read

		_, err = db.Update(replyObj)
		checkErr(err, "ERROR UpdateReplySent ERROR")
	}

	return http.StatusOK, ""

}
