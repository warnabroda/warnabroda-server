package routes

import (
	"os"
	"net/http"
	"strconv"
	"math/rand"
	"strings"
	"io/ioutil"	
	"time"
//	"fmt"	

	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"bitbucket.org/hbtsmith/warnabrodagomartini/messages"
	"github.com/go-martini/martini"
	"github.com/coopernurse/gorp"
)

const (	
	SQL_IN_IGNORE_LIST_BY_CONTACT			= "SELECT * FROM ignore_list WHERE Contact= :contact"
	SQL_IN_IGNORE_LIST_BY_CODE				= "SELECT * FROM ignore_list WHERE confirmation_code= :code"
	SQL_COUNT_MULTIPLE_IGNOREME_REQUESTS	= "SELECT COUNT(*) FROM ignore_list WHERE ip=? AND (created_date + INTERVAL 2 HOUR) > NOW()"
)



// Generate random A-Z letters 6 sized for ignore list confirmation purpose
func randomString(l int ) string {
    bytes := make([]byte, l)
    for i:=0 ; i<l ; i++ {
        bytes[i] = byte(randInt(65,90))
    }
    return string(bytes)
}

// Generate random number based upon a min and max range
func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

// opens the template, parse the variables sets the email struct and Send the confirmation code to confirm the ignored contact.
func sendEmailIgnoreme(entity *models.Ignore_List, db gorp.SqlExecutor){
	//reads the e-mail template from a local file
	wab_email_template := wab_root + "/resource/ignoreme_"+entity.Lang_key+".html"
	template_byte, err := ioutil.ReadFile(wab_email_template)
	checkErr(err, "Ignore-me Email File Opening ERROR")
	template_email_string := string(template_byte[:])
	
	var email_content string
	email_content = strings.Replace(template_email_string, "{{code}}", entity.Confirmation_code, 1)
	email_content = strings.Replace(email_content, "{{url_ignoreme}}", models.URL_IGNOREME, 2)
	email_content = strings.Replace(email_content, "{{url_contacus}}", models.URL_CONTACT_US, 1)
	email_content = strings.Replace(email_content, "{{email}}", models.EMAIL_WARNABRODA, 2)

	email := &models.Email{
		TemplatePath: wab_email_template,	
		Content: email_content, 	
		Subject: messages.GetLocaleMessage(entity.Lang_key,"MSG_EMAIL_SUBJECT_ADD_IGNORE_LIST"),		
		ToAddress: entity.Contact,
		FromName: models.WARN_A_BRODA,
		LangKey: entity.Lang_key,
		Async: false,
		UseContent: true,
		HTMLContent: true,
	}	
	
	sent, response := SendMail(email)
	if sent {
		entity.Message = response
		UpdateIgnoreList(entity, db)
	}
}

// send a SMS with the confirmation code to confirm the ignored contact
func sendSMSIgnoreme(entity *models.Ignore_List, db gorp.SqlExecutor){	

	sms := &models.SMS {
		CredencialKey: os.Getenv("WARNACREDENCIAL"),  
	    Content: strings.Replace(messages.GetLocaleMessage(entity.Lang_key,"MSG_SMS_IGNORE_CONFIRMATION_REQUEST"), "{{url}}", entity.Confirmation_code, 1),
	    URLPath: models.URL_MAIN_MOBILE_PRONTO,	  
	    Scheme: "http",	  
	    Host: models.URL_DOMAIN_MOBILE_PRONTO,	  
	    Project: os.Getenv("WARNAPROJECT"),	  
	    AuxUser: "WAB",	      
	    MobileNumber: strings.Replace(entity.Contact, "+", "", 1),
	    SendProject:"N",	    
	}

	sent, response := SendSMS(sms)

	if  sent {		
		entity.Message = response
		UpdateIgnoreList(entity, db)
	}
}

// Add the request to be ignored for future warnings, it requires further confimation
func AddIgnoreList(entity models.Ignore_List, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {
	
	status := &models.DefaultStruct{
			Id:       http.StatusOK,
			Name:     messages.GetLocaleMessage(entity.Lang_key,"MSG_CONFIRM_IGNOREME"),
			Lang_key: entity.Lang_key,
		}

	if MoreThanTwoRequestByIp(db, &entity){
		status = &models.DefaultStruct{
			Id:       http.StatusForbidden,
			Name:     messages.GetLocaleMessage(entity.Lang_key,"MSG_TOO_MANY_IGNOREME_REQUESTS"),
			Lang_key: entity.Lang_key,
		}
		return http.StatusCreated, Must(enc.EncodeOne(status))		
	} 

	ingnored := InIgnoreList(db, entity.Contact)	

	if ingnored!= nil && ingnored.Confirmed {

		status = &models.DefaultStruct{
			Id:       http.StatusForbidden,
			Name:     messages.GetLocaleMessage(entity.Lang_key,"MSG_CONTACT_ALREADY_IGNORED"),
			Lang_key: entity.Lang_key,
		}

		return http.StatusCreated, Must(enc.EncodeOne(status))

	} else if ingnored!= nil {
		status = &models.DefaultStruct{
			Id:       http.StatusForbidden,
			Name:     messages.GetLocaleMessage(entity.Lang_key, "MSG_IGNORE_REQUEST_EXISTS"),
			Lang_key: entity.Lang_key,
		}

		return http.StatusCreated, Must(enc.EncodeOne(status))
	}
	
    rand.Seed(time.Now().UTC().UnixNano())   
	entity.Created_by 			= "user"	
	entity.Confirmed 			= false;
	entity.Confirmation_code 	= randomString(6)
	
	errIns := db.Insert(&entity)
	checkErr(errIns, "INSERT IGNORE FAIL")

	if strings.Contains(entity.Contact, "@"){
		status.Name += " via e-mail."			
		go sendEmailIgnoreme(&entity, db)
	} else if strings.Contains(entity.Contact, "+55"){
		status.Name += " via SMS."
		go sendSMSIgnoreme(&entity, db)
	} 

	
		
	//w.Header().Set("Location", fmt.Sprintf("/warnabroda/ignore-list/%d", entity.Id))
	return http.StatusCreated, Must(enc.EncodeOne(status))
}

// Confirm the request for ignore list
func ConfirmIgnoreList(entity models.Ignore_List, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {
	
	status := &models.DefaultStruct{
			Id:       http.StatusOK,
			Name:     messages.GetLocaleMessage(entity.Lang_key, "MSG_IGNORED_SUCCESSFUL"),
			Lang_key: entity.Lang_key,
		}

	ignored := GetIgnoreContact(db, entity.Confirmation_code)

	if ignored != nil {
		ignored.Confirmed = true
		ignored.Last_modified_date = entity.Last_modified_date
		UpdateIgnoreList(ignored, db)
	} else {
		status = &models.DefaultStruct{
			Id:       http.StatusForbidden,
			Name:     messages.GetLocaleMessage(entity.Lang_key, "MSG_IGNOREME_CODE_INVALID"),
			Lang_key: entity.Lang_key,
		}
	}

	return http.StatusCreated, Must(enc.EncodeOne(status))
}

func DeleteIgnoreList(db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	obj, _ := db.Get(models.Ignore_List{}, id)
	if err != nil || obj == nil {
		checkErr(err, "get failed")
		// Invalid id, or does not exist
		return http.StatusNotFound, ""
	}
	entity := obj.(*models.Ignore_List)
	_, err = db.Delete(entity)
	if err != nil {
		checkErr(err, "delete failed")
		return http.StatusConflict, ""
	}
	return http.StatusNoContent, ""
}

func UpdateIgnoreList(entity *models.Ignore_List, db gorp.SqlExecutor) {
		
	_, err := db.Update(entity)
	if err != nil {
		checkErr(err, "update failed")
	}
}

// Check if the contact already requested an ignore list add.
// In case the contact exists on the list the method returns it
func InIgnoreList(db gorp.SqlExecutor, contact string) *models.Ignore_List {	

	ignored := models.Ignore_List{}

	err := db.SelectOne(&ignored, SQL_IN_IGNORE_LIST_BY_CONTACT,  
		map[string]interface{}{
	  		"contact": contact,
		})	
	if err != nil {
		return nil
	}	


	return &ignored
}

// Get an existent ignoreme register, in case there is none returns nil
func GetIgnoreContact(db gorp.SqlExecutor, id string) *models.Ignore_List {	

	var ignored models.Ignore_List
	err := db.SelectOne(&ignored, SQL_IN_IGNORE_LIST_BY_CODE, 
		map[string]interface{}{
	  		"code": id, 	  		
		})
	if err != nil {
		return nil
	}
	
	return &ignored
}

func GetIgnoreContactById(id int64, db gorp.SqlExecutor) *models.Ignore_List {
	obj, err := db.Get(models.Ignore_List{}, id)

	if err != nil || obj == nil {	
		return nil
	}

	entity := obj.(*models.Ignore_List)
	return entity
}

// intercepts more than two requests to ignore list add.
func MoreThanTwoRequestByIp(db gorp.SqlExecutor, entity *models.Ignore_List) bool{

	sql := SQL_COUNT_MULTIPLE_IGNOREME_REQUESTS

	total, err := db.SelectInt(sql, entity.Ip)
	checkErr(err, "COUNT ERROR")

	return total >= 2

}

func UpdateIgnoreSent(entity *models.Ignore_List, db gorp.SqlExecutor) bool {
	entity.Sent = true
	entity.Last_modified_date = time.Now().String()
	_, err := db.Update(entity)
	checkErr(err, "ERROR UpdateWarningSent ERROR")	
	return err == nil
}