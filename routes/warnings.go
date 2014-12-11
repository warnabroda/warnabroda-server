package routes

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
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
	wab_root 			= os.Getenv("WARNAROOT")		
)

func sendEmailWarn(entity *models.Warning, db gorp.SqlExecutor) {	

	//reads the e-mail template from a local file
	wab_email_template := wab_root + "/models/warning.html"
	template_byte, err := ioutil.ReadFile(wab_email_template)
	checkErr(err, "File Opening ERROR")
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
		FromName: "Warn A Broda",
		LangKey: "br",
		Async: false,
		UseContent: true,
		HTMLContent: true,
	}	
	
	if SendMail(email, db) {
		UpdateWarningSent(entity, db)
	} else {
		fmt.Println("ERROR SENDING MAIL")
	}

}

func sendSMSWarn(entity *models.Warning, db gorp.SqlExecutor){

	message := SelectMessage(db, entity.Id_message)
	sms_message := "Ola Broda, "
	sms_message += "Você " + message.Name + ". "
	sms_message += "Warn A Broda você também: www.warnabroda.com"

	sms := &models.SMS {
		CredencialKey: os.Getenv("WARNACREDENCIAL"),  
	    Content: sms_message,
	    URLPath: "http://www.mpgateway.com/v_2_00/smsfollow/smsfollow.aspx?",	  
	    Scheme: "http",	  
	    Host: "www.mpgateway.com",	  
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

func GetWarnings(enc Encoder, db gorp.SqlExecutor) (int, string) {
	
	var warnings []models.Warning
	_, err := db.Select(&warnings, "SELECT * FROM warnings ORDER BY id")
	checkErr(err, "select failed")

	if err != nil {
		return http.StatusInternalServerError, ""
	}
	return http.StatusOK, Must(enc.Encode(warningsToIface(warnings)...))
}

func CountWarnings(enc Encoder, db gorp.SqlExecutor) (int, string) {
	
	total, err := db.SelectInt("SELECT COUNT(*) AS total FROM warnings WHERE Sent=true")
	checkErr(err, "COUNT ERROR")
	
	if err != nil {
		return http.StatusInternalServerError, ""
	}
	return http.StatusOK, strconv.FormatInt(total, 10)
}

func GetWarning(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	checkErr(err, "get failed")

	obj, _ := db.Get(models.Warning{}, id)

	if err != nil || obj == nil {	
		return http.StatusNotFound, ""
	}

	entity := obj.(*models.Warning)
	return http.StatusOK, Must(enc.EncodeOne(entity))
}

func AddWarning(entity models.Warning, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {

	status := &models.Message{
		Id:       200,
		Name:     "Broda Avisado(a)",
		Lang_key: "br",
	}

	entity.Sent = false
	entity.Created_by = "system"
	entity.Created_date = time.Now().String()
	entity.Lang_key = "br"

	err := db.Insert(&entity)
	checkErr(err, "insert failed")
	if err != nil {
		return http.StatusConflict, ""
	}
	w.Header().Set("Location", fmt.Sprintf("/warnabroda/warnings/%d", entity.Id))

	ingnored := InIgnoreList(db, entity.Contact)

	if ingnored!= nil && ingnored.Confirmed {
		status = &models.Message{
			Id:       403,
			Name:     "Este Broda está na Ignore List, pois não deseja receber avisos do Warn A Broda, Sorry ae.",
			Lang_key: "br",
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

func processEmail(warning *models.Warning, db gorp.SqlExecutor, status *models.Message){
	
	if emailSentToContact(warning, db) {
		status.Id = 403
		status.Name = "Broda já foi avisado(a) há instantes atrás. Muito Obrigado."
		status.Lang_key = "br"
	} else {
		go sendEmailWarn(warning, db)
	}
}

func processSMS(warning *models.Warning, db gorp.SqlExecutor, status *models.Message) {

	if smsSentToContact(warning, db) {
		status.Id = 403
		status.Name = "Este número já recebeu um SMS hoje ou seu IP(" + warning.Ip + ") já enviou a cota maxima de SMS diário."
		status.Lang_key = "br"
	} else {
		go sendSMSWarn(warning, db)
	}

}

func emailSentToContact(warning *models.Warning, db gorp.SqlExecutor) bool {
	
	now_lower 		:= time.Now().Add(-2*time.Hour)
	now_upper 		:= time.Now().Add(2*time.Hour)
	str_now_lower	:= now_lower.Format(models.DbFormat)			
	str_now_upper	:= now_upper.Format(models.DbFormat)	
		
	select_statement := " SELECT COUNT(*) FROM warnings "
	select_statement += " WHERE Id_contact_type = 1 AND Sent = 1 AND "
	select_statement += " Contact = '" + warning.Contact + "' AND "
	select_statement += " Created_date BETWEEN '" + str_now_lower +"' AND '" + str_now_upper +"' AND "
	select_statement += " Id <> " + strconv.FormatInt(warning.Id, 10)
	
	exists, err 	:= db.SelectInt(select_statement)
	checkErr(err, "Checking Contact failed")
	
	return exists > 0
}

func smsSentToContact(warning *models.Warning, db gorp.SqlExecutor) bool {

	str_today 		:= time.Now().Format(models.JsonFormat)

	select_statement := " SELECT COUNT(*) FROM warnings "
	select_statement += " WHERE Id_contact_type = 2 AND Sent = 1 AND "
	select_statement += " (Contact = '" + warning.Contact + "' OR Ip LIKE '%" + warning.Ip + "%' ) AND "
	select_statement += " Created_date BETWEEN '" + str_today + " 00:00:00' AND '" + str_today + " 23:59:59' AND "
	select_statement += " Id <> " + strconv.FormatInt(warning.Id, 10)

	exists, err 	:= db.SelectInt(select_statement)
	checkErr(err, "Checking Contact failed")
	
	return exists > 0
}

func UpdateWarningSent(entity *models.Warning, db gorp.SqlExecutor) {
	entity.Sent = true
	entity.Last_modified_date = time.Now().String()
	_, err := db.Update(entity)
	checkErr(err, "update failed")	
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
