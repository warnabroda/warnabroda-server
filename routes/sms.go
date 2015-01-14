package routes

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"bitbucket.org/hbtsmith/warnabrodagomartini/messages"	
	"github.com/coopernurse/gorp"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"os"
	"fmt"
)

func ProcessSMS(warning *models.Warning, db gorp.SqlExecutor, status *models.DefaultStruct) {
	
	
	if isWarnSentLimitByIpOver(warning, db) {
		status.Id = http.StatusForbidden
		status.Name = strings.Replace(messages.GetLocaleMessage(warning.Lang_key,"MSG_SMS_QUOTA_EXCEEDED"), "{{ip}}", warning.Ip, 1) 
		status.Lang_key = warning.Lang_key
	} else {
		go sendSMSWarn(warning, db)
	}
}

func isWarnSentLimitByIpOver(warning *models.Warning, db gorp.SqlExecutor) bool{
	
	exists, err 	:= db.SelectInt(BuildCountWarningsSql("ip"), map[string]interface{}{
		"id_contact_type": warning.Id_contact_type,
		"sent": true,
		"interval": 24,
		"ip": warning.Ip,
		})
	checkErr(err, "SELECT isWarnSentLimitByIpOver ERROR")
	fmt.Println(exists)
	return exists > 3
}

func sendSMSWarn(entity *models.Warning, db gorp.SqlExecutor){

	message := SelectMessage(db, entity.Id_message)
	sms_message := strings.Replace(messages.GetLocaleMessage(entity.Lang_key,"MSG_SMS_BODY"), "{{body}}", message.Name, 1)	
	sms_message += messages.GetLocaleMessage(entity.Lang_key,"MSG_SMS_FOOTER")

	sms := &models.SMS {
		CredencialKey: os.Getenv("WARNACREDENCIAL"),  
	    Content: sms_message,
	    URLPath: models.URL_MAIN_MOBILE_PRONTO,	  
	    Scheme: "http",	  
	    Host: models.URL_DOMAIN_MOBILE_PRONTO,	  
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

// Component to send a SMS using mobile pronto
func SendSMS(sms *models.SMS, db gorp.SqlExecutor) (bool, string) {
	u, err := url.Parse(sms.URLPath)
	checkErr(err, "Ugly URL")	

	u.Scheme = sms.Scheme
	u.Host = sms.Host
	q := u.Query()
	q.Set("CREDENCIAL", sms.CredencialKey)
	q.Set("PRINCIPAL_USER", sms.Project)
	q.Set("AUX_USER", sms.AuxUser)
	q.Set("MOBILE", sms.MobileNumber)
	q.Set("SEND_PROJECT", sms.SendProject)
	q.Set("MESSAGE", sms.Content)
	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())	
	 checkErr(err, "SMS Not Sent")

	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	 checkErr(err, "No response from SMS Sender")

	//return true, "TESTE"
	return err == nil, string(robots[:])
}