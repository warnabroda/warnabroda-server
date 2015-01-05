package routes

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"bitbucket.org/hbtsmith/warnabrodagomartini/i18n"
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"	
//	"io/ioutil"
	"net/http"
//	"os"
	"strconv"
	"strings"
	"time"
)

const (
	
	MSG_SMS_HEADER				= "Ola Broda, "
	MSG_SMS_BODY				= "Você {{body}}."
	MSG_SMS_FOOTER				= "Warn A Broda você também: www.warnabroda.com"
	MSG_WARNING_SENT_SUCCESS 	= "Broda Avisado(a)"
	MSG_IGNORED_USER			= "Este Broda está na Ignore List, pois não deseja receber avisos do Warn A Broda, Sorry ae."	
	MSG_SMS_QUOTA_EXCEEDED		= "Este IP( {{ip}} ) já enviou a cota maxima de SMS diário."
	MSG_SMS_SAME_WARN_BY_IP		= "Este IP( {{ip}} ) já enviou esta mesma mensagem para este Broda há menos de {{time}} horas."
	MSG_SMS_SAME_WARN_DIFF_IP	= "Este Broda recebeu esta mesma mensagem 2 ou mais vezes há menos de {{time}} horas."
	SQL_WARNING_BYID			= "SELECT * FROM warnings ORDER BY id"
	SQL_CHECK_SENT_WARN			= " SELECT COUNT(*) FROM warnings " + 
							  	" WHERE Id_contact_type = :id_contact_type AND Sent = true AND " + 							  	
							  	" (Contact = :contact OR Ip LIKE :ip ) AND " +
							  	" Created_date BETWEEN :lower_str_date AND :upper_str_date AND " + 
							  	" Id <> :id "
)

func BuildCountWarningsSql(count_by string) string {

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

func GetWarning(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	checkErr(err, "GET MARTINI PARAM ERROR")

	obj, _ := db.Get(models.Warning{}, id)

	if err != nil || obj == nil {	
		return http.StatusNotFound, ""
	}

	entity := obj.(*models.Warning)
	return http.StatusOK, Must(enc.EncodeOne(entity))
}

func UpdateWarningSent(entity *models.Warning, db gorp.SqlExecutor) {
	entity.Sent = true
	entity.Last_modified_date = time.Now().String()
	_, err := db.Update(entity)
	checkErr(err, "ERROR UpdateWarningSent ERROR")	
}

func AddWarning(entity models.Warning, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {

	status := &models.DefaultStruct{
		Id:       http.StatusOK,
		Name:     MSG_WARNING_SENT_SUCCESS,
		Lang_key: i18n.BR_LANG_KEY,
	}

	entity.Sent = false
	entity.Created_by = "system"
	entity.Created_date = time.Now().String()
	entity.Lang_key = i18n.BR_LANG_KEY

	err := db.Insert(&entity)
	checkErr(err, "INSERT WARNING ERROR")
	if err != nil {
		return http.StatusConflict, ""
	}
	w.Header().Set("Location", fmt.Sprintf("/warnabroda/warnings/%d", entity.Id))

	ingnored := InIgnoreList(db, entity.Contact)

	if ingnored!= nil && ingnored.Confirmed {
		status = &models.DefaultStruct{
			Id:       http.StatusForbidden,
			Name:     MSG_IGNORED_USER,
			Lang_key: i18n.BR_LANG_KEY,
		}
	} else {
		processWarn(&entity, db, status)
	}

	return http.StatusCreated, Must(enc.EncodeOne(status))
}

func processWarn(warning *models.Warning, db gorp.SqlExecutor, status *models.DefaultStruct){

	if isSameWarnSentByIp(warning, db) {
		status.Id = http.StatusForbidden
		status.Name = strings.Replace(MSG_SMS_SAME_WARN_BY_IP, "{{ip}}", warning.Ip, 1) 
		status.Name = strings.Replace(status.Name, "{{time}}", "2", 1)
		status.Lang_key = i18n.BR_LANG_KEY
	} else if isSameWarnSentTwiceOrMoreDifferentIp(warning, db) {
		status.Id = http.StatusForbidden
		status.Name = strings.Replace(MSG_SMS_SAME_WARN_DIFF_IP, "{{time}}", "2", 1) 		
		status.Lang_key = i18n.BR_LANG_KEY
	} else {
		if warning.Id_contact_type == 1 {
			ProcessEmail(warning, db)
		} else if warning.Id_contact_type == 2 {
			ProcessSMS(warning, db, status)
		}
	}
}

func isSameWarnSentByIp(warning *models.Warning, db gorp.SqlExecutor) bool {		

	exists, err 	:= db.SelectInt(BuildCountWarningsSql("same_message_by_ip"), map[string]interface{}{		
		"sent": true,
		"contact": warning.Contact,
		"interval": 2,
		"id_message": warning.Id_message,
		"ip": warning.Ip,
		})
	checkErr(err, "SELECT isSameWarnSentByIp ERROR")
	
	return exists > 0
}

func isSameWarnSentTwiceOrMoreDifferentIp(warning *models.Warning, db gorp.SqlExecutor) bool {		

	exists, err 	:= db.SelectInt(BuildCountWarningsSql("same_message"), map[string]interface{}{		
		"sent": true,
		"contact": warning.Contact,
		"interval": 2,
		"id_message": warning.Id_message,		
		})
	checkErr(err, "SELECT isSameWarnSentTwiceOrMoreDifferentIp ERROR")
	
	return exists > 1
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
