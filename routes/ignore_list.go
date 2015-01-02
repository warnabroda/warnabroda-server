package routes

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"bitbucket.org/hbtsmith/warnabrodagomartini/i18n"
	"github.com/go-martini/martini"
	"github.com/coopernurse/gorp"
	"os"
	"net/http"
	"strconv"
	"time"
	"math/rand"
	"strings"
	"io/ioutil"	
//	"fmt"	
)

const (
	URL_IGNOREME 							= "www.warnabroda.com/#/ignoreme"
	MSG_IGNOREME_CODE_INVALID				= "Código inválido."
	MSG_IGNORED_SUCCESSFUL					= "Ignorado com Sucesso, se um dia você se arrepender, entre em contato conosco é a unica forma de voltar a participar do Warn A Broda."
	MSG_CONTACT_ALREADY_IGNORED				= "Contato Já estava na Lista de Ignorados!"
	MSG_IGNORE_REQUEST_EXISTS				= "Solicitações de bloqeuio expiram em 24 horas. Aguarde para solicitar novamente ou entre em contato com o Warn A Broda."
	MSG_TOO_MANY_IGNOREME_REQUESTS			= "Ooopa! Pra que tantas solicitações Broda? Se você possui mais de 2 contatos à bloquear por vez, entre em contato com o Warn A Broda."
	MSG_CONFIRM_IGNOREME					= "Favor confirmar bloqueio de contato"
	MSG_EMAIL_SUBJECT_ADD_IGNORE_LIST		= "Adicionar contato à ignore list do Warn A Broda"
	MSG_SMS_IGNORE_CONFIRMATION_REQUEST		= "Pro Warn A Broda lhe ignorar efetivamente, " +
											"por favor entre em: " + 
											URL_IGNOREME + 
											" e informe o codigo: "
	SQL_IN_IGNORE_LIST_BY_CONTACT			= "SELECT * FROM ignore_list WHERE Contact= :contact "
	SQL_IN_IGNORE_LIST_BY_CODE				= "SELECT * FROM ignore_list WHERE confirmation_code= :code"
	SQL_REMOVE_OLD_IGNOREME_REQUESTS		= "DELETE FROM ignore_list WHERE confirmed = false AND (created_date + INTERVAL 24 HOUR) < NOW()"
	SQL_COUNT_MULTIPLE_IGNOREME_REQUESTS	= "SELECT COUNT(*) FROM ignore_list WHERE ip=? AND (created_date + INTERVAL 2 HOUR) > NOW()"
)

//Initialize all required functions when container is up.
func init(){

	IgnoreListCleaner()	
}

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
	wab_email_template := wab_root + "/models/ignoreme.html"
	template_byte, err := ioutil.ReadFile(wab_email_template)
	checkErr(err, "Ignore-me Email File Opening ERROR")
	template_email_string := string(template_byte[:])
	
	var email_content string
	email_content = strings.Replace(template_email_string, "{{code}}", entity.Confirmation_code, 1)
	email_content = strings.Replace(email_content, "{{url}}", URL_IGNOREME, 2)

	email := &models.Email{
		TemplatePath: wab_email_template,	
		Content: email_content, 	
		Subject: MSG_EMAIL_SUBJECT_ADD_IGNORE_LIST,		
		ToAddress: entity.Contact,
		FromName: i18n.WARN_A_BRODA,
		LangKey: i18n.BR_LANG_KEY,
		Async: false,
		UseContent: true,
		HTMLContent: true,
	}	
	
	SendMail(email, db)
}

// send a SMS with the confirmation code to confirm the ignored contact
func sendSMSIgnoreme(entity *models.Ignore_List, db gorp.SqlExecutor){	

	sms := &models.SMS {
		CredencialKey: os.Getenv("WARNACREDENCIAL"),  
	    Content: MSG_SMS_IGNORE_CONFIRMATION_REQUEST + entity.Confirmation_code,
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
		UpdateIgnoreList(entity, db)	
	}
}

// Add the request to be ignored for future warnings, it requires further confimation
func AddIgnoreList(entity models.Ignore_List, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {
	
	status := &models.DefaultStruct{
			Id:       http.StatusOK,
			Name:     MSG_CONFIRM_IGNOREME,
			Lang_key: i18n.BR_LANG_KEY,
		}

	if MoreThanTwoRequestByIp(db, &entity){
		status = &models.DefaultStruct{
			Id:       http.StatusForbidden,
			Name:     MSG_TOO_MANY_IGNOREME_REQUESTS,
			Lang_key: i18n.BR_LANG_KEY,
		}
		return http.StatusCreated, Must(enc.EncodeOne(status))		
	} 

	ingnored := InIgnoreList(db, entity.Contact)	

	if ingnored!= nil && ingnored.Confirmed {

		status = &models.DefaultStruct{
			Id:       http.StatusForbidden,
			Name:     MSG_CONTACT_ALREADY_IGNORED,
			Lang_key: i18n.BR_LANG_KEY,
		}

		return http.StatusCreated, Must(enc.EncodeOne(status))

	} else if ingnored!= nil {
		status = &models.DefaultStruct{
			Id:       http.StatusForbidden,
			Name:     MSG_IGNORE_REQUEST_EXISTS,
			Lang_key: i18n.BR_LANG_KEY,
		}

		return http.StatusCreated, Must(enc.EncodeOne(status))
	}
	
    rand.Seed(time.Now().UTC().UnixNano())   
	entity.Created_by 			= "user"
	entity.Created_date 		= time.Now().String()	
	entity.Confirmed 			= false;
	entity.Confirmation_code 	= randomString(6)

	if strings.Contains(entity.Contact,"@"){
		status.Name += " via e-mail."			
		go sendEmailIgnoreme(&entity, db)
	} else {
		status.Name += " via SMS."
		go sendSMSIgnoreme(&entity, db)
	}		

	errIns := db.Insert(&entity)
	checkErr(errIns, "INSERT IGNORE FAIL")
	
		
	//w.Header().Set("Location", fmt.Sprintf("/warnabroda/ignore-list/%d", entity.Id))
	return http.StatusCreated, Must(enc.EncodeOne(status))
}

// Confirm the request for ignore list
func ConfirmIgnoreList(entity models.Ignore_List, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {
	
	status := &models.DefaultStruct{
			Id:       http.StatusOK,
			Name:     MSG_IGNORED_SUCCESSFUL,
			Lang_key: i18n.BR_LANG_KEY,
		}

	ignored := GetIgnoreContact(db, entity.Confirmation_code)

	if ignored != nil {
		ignored.Confirmed = true
		UpdateIgnoreList(ignored, db)
	} else {
		status = &models.DefaultStruct{
			Id:       http.StatusForbidden,
			Name:     MSG_IGNOREME_CODE_INVALID,
			Lang_key: i18n.BR_LANG_KEY,
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
	
	entity.Last_modified_date = time.Now().String()
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

// Remove all requests non confirmed older than 24 hours
func IgnoreListCleaner(){
	sql := SQL_REMOVE_OLD_IGNOREME_REQUESTS
	models.Dbm.Exec(sql)
	//fmt.Println(sql)
	ticker := time.NewTicker(time.Hour)
	quit := make(chan struct{})
	go func() {
	    for {
	       select {
	        case <- ticker.C:	        	
	            models.Dbm.Exec(sql)
	        	//fmt.Println(sql)
	        case <- quit:
	            ticker.Stop()
	            return
	        }
	    }
	 }()
}

// intercepts more than two requests to ignore list add.
func MoreThanTwoRequestByIp(db gorp.SqlExecutor, entity *models.Ignore_List) bool{

	sql := SQL_COUNT_MULTIPLE_IGNOREME_REQUESTS

	total, err := db.SelectInt(sql, entity.Ip)
	checkErr(err, "COUNT ERROR")

	return total >= 2

}