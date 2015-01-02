package routes

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"bitbucket.org/hbtsmith/warnabrodagomartini/i18n"
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"	
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	wab_root 				= os.Getenv("WARNAROOT")
)

const (
	
	MSG_SMS_HEADER				= "Ola Broda, "
	MSG_SMS_BODY				= "Você {{body}}."
	MSG_SMS_FOOTER				= "Warn A Broda você também: www.warnabroda.com"
	MSG_WARNING_SENT_SUCCESS 	= "Broda Avisado(a)"
	MSG_IGNORED_USER			= "Este Broda está na Ignore List, pois não deseja receber avisos do Warn A Broda, Sorry ae."
	MSG_EMAIL_WARNING_SENT		= "Broda já foi avisado(a) há instantes atrás. Muito Obrigado."	
	MSG_SMS_QUOTA_EXCEEDED		= "Este número já recebeu um SMS hoje ou seu IP( {{ip}} ) já enviou a cota maxima de SMS diário."
	SQL_WARNING_BYID			= "SELECT * FROM warnings ORDER BY id"
	SQL_CHECK_SENT_WARN			= " SELECT COUNT(*) FROM warnings " + 
							  	" WHERE Id_contact_type = :id_contact_type AND Sent = true AND " + 
							  	" Contact = :contact  AND "  + 
							  	" Created_date BETWEEN :lower_str_date AND :upper_str_date AND " + 
							  	" Id <> :id "
)

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
		if entity.Id_contact_type == 1 {
			processEmail(&entity, db, status)
		} else if entity.Id_contact_type == 2 {
			processSMS(&entity, db, status)
		}
	}

	return http.StatusCreated, Must(enc.EncodeOne(status))
}

func processEmail(warning *models.Warning, db gorp.SqlExecutor, status *models.DefaultStruct){	

	if emailSentToContact(warning, db) {
		status.Id = http.StatusForbidden
		status.Name = MSG_EMAIL_WARNING_SENT
		status.Lang_key = i18n.BR_LANG_KEY
	} else {
		go sendEmailWarn(warning, db)
	}
}

func emailSentToContact(warning *models.Warning, db gorp.SqlExecutor) bool {
	
	now_lower 		:= time.Now().Add(-1*time.Hour)
	now_upper 		:= time.Now().Add(1*time.Hour)
	
	fmt.Println(SQL_CHECK_SENT_WARN)	
	exists, err 	:= db.SelectInt(SQL_CHECK_SENT_WARN, map[string]interface{}{
		"id_contact_type": 1,
		"contact": warning.Contact,
		"lower_str_date": now_lower.Format(models.DbFormat),
		"upper_str_date": now_upper.Format(models.DbFormat),
		"id": strconv.FormatInt(warning.Id, 10),
		})
	checkErr(err, "SELECT ALREADY EMAILED CONTACT ERROR")
	
	return exists > 0
}

func sendEmailWarn(entity *models.Warning, db gorp.SqlExecutor) {	

	//reads the e-mail template from a local file
	wab_email_template := wab_root + "/models/warning.html"
	template_byte, err := ioutil.ReadFile(wab_email_template)
	checkErr(err, "Email File Opening ERROR")
	template_email_string := string(template_byte[:])

	subject := GetRandomSubject()
	message := SelectMessage(db, entity.Id_message)
	var email_content string
	email_content = strings.Replace(template_email_string, "{{warning}}", message.Name, 1)

	email := &models.Email{
		TemplatePath: wab_email_template,	
		Content: email_content, 	
		Subject: subject.Name,		
		ToAddress: entity.Contact,
		FromName: i18n.WARN_A_BRODA,
		LangKey: i18n.BR_LANG_KEY,
		Async: false,
		UseContent: true,
		HTMLContent: true,
	}	
	
	if SendMail(email, db) {
		UpdateWarningSent(entity, db)
	} else {
		fmt.Println("SENDING MAIL ERROR")
	}

}

func processSMS(warning *models.Warning, db gorp.SqlExecutor, status *models.DefaultStruct) {
	
	if smsSentToContact(warning, db) {
		status.Id = http.StatusForbidden
		status.Name = strings.Replace(MSG_SMS_QUOTA_EXCEEDED, "{{ip}}", warning.Ip, 1) 
		status.Lang_key = i18n.BR_LANG_KEY
	} else {
		go sendSMSWarn(warning, db)
	}
}

func smsSentToContact(warning *models.Warning, db gorp.SqlExecutor) bool {		

	exists, err 	:= db.SelectInt(SQL_CHECK_SENT_WARN, map[string]interface{}{
		"id_contact_type": 2,
		"contact": warning.Contact,
		"lower_str_date": time.Now().Format(models.JsonFormat) + " 00:00:00",
		"upper_str_date": time.Now().Format(models.JsonFormat) + " 23:59:59",
		"id": strconv.FormatInt(warning.Id, 10),
		})
	checkErr(err, "SELECT ALREADY SMSED CONTACT ERROR")
	
	return exists > 0
}

func sendSMSWarn(entity *models.Warning, db gorp.SqlExecutor){

	message := SelectMessage(db, entity.Id_message)
	sms_message := MSG_SMS_HEADER
	sms_message += strings.Replace(MSG_SMS_BODY, "{{body}}", message.Name, 1)
	sms_message += MSG_SMS_FOOTER

	sms := &models.SMS {
		CredencialKey: os.Getenv("WARNACREDENCIAL"),  
	    Content: sms_message,
	    URLPath: i18n.URL_MAIN_MOBILE_PRONTO,	  
	    Scheme: "http",	  
	    Host: i18n.URL_DOMAIN_MOBILE_PRONTO,	  
	    Project: os.Getenv("WARNAPROJECT"),	  
	    AuxUser: "WAB",	      
	    MobileNumber: "55"+entity.Contact,
	    SendProject:"N",	    
	}

	sent, response := SendSMS(sms, db)

	if  sent {
		entity.Message = response
		UpdateWarningSent(entity, db)
	}
}

func UpdateWarningSent(entity *models.Warning, db gorp.SqlExecutor) {
	entity.Sent = true
	entity.Last_modified_date = time.Now().String()
	_, err := db.Update(entity)
	checkErr(err, "WARNING UPDATE ERROR")	
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
