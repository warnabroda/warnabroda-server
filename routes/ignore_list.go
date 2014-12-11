package routes

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"os"
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
	"time"
	"math/rand"
	"strings"
	"io/ioutil"
)

const (
	ignore_url = "www.warnabroda.com/#/ignoreme"
)

func randomString(l int ) string {
    bytes := make([]byte, l)
    for i:=0 ; i<l ; i++ {
        bytes[i] = byte(randInt(65,90))
    }
    return string(bytes)
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

func GetIgnoreContact(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	
	id, err := strconv.Atoi(parms["id"])

	var ignored models.Ignore_List
	err = db.SelectOne(&ignored, "SELECT * FROM ignore_list WHERE confirmation_code=?", id)
	if err != nil {
		checkErr(err, "select failed")
		return http.StatusInternalServerError, ""
	}

	UpdateIgnoreList(&ignored, db)
	
	return http.StatusOK, Must(enc.EncodeOne(ignored))
}

func sendEmailIgnoreme(entity *models.Ignore_List, db gorp.SqlExecutor){
	//reads the e-mail template from a local file
	wab_email_template := wab_root + "/models/ignoreme.html"
	template_byte, err := ioutil.ReadFile(wab_email_template)
	checkErr(err, "File Opening ERROR")
	template_email_string := string(template_byte[:])
	
	var email_content string
	email_content = strings.Replace(template_email_string, "{{code}}", entity.Confirmation_code, 1)
	email_content = strings.Replace(email_content, "{{url}}", ignore_url, 2)

	email := &models.Email{
		TemplatePath: wab_email_template,	
		Content: email_content, 	
		Subject: "Adicionar contato à ignore list do Warn A Broda",		
		ToAddress: entity.Contact,
		FromName: "Warn A Broda",
		LangKey: "br",
		Async: false,
		UseContent: true,
		HTMLContent: true,
	}	
	
	SendMail(email, db)
}

func sendSMSIgnoreme(entity *models.Ignore_List, db gorp.SqlExecutor){
	
	sms_message := "Pro Warn A Broda lhe ignorar efetivamente, "
	sms_message += "por favor entre em: "
	sms_message += "www.warnabroda.com/#/ignoreme "
	sms_message += "e informe o codigo: "+entity.Confirmation_code

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
		UpdateIgnoreList(entity, db)	
	}
}

func AddIgnoreList(entity models.Ignore_List, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {
	
	status := &models.Message{
		Id:       200,
		Name:     "Favor confirmar bloqueio de contato",
		Lang_key: "br",
	}	

	ingnored := InIgnoreList(db, entity.Contact)

	if ingnored!=nil && ingnored.Confirmed {

		status = &models.Message{
			Id:       403,
			Name:     "Contato Já estava na Lista de Ignorados!",
			Lang_key: "br",
		}

		return http.StatusCreated, Must(enc.EncodeOne(status))

	} else if ingnored!=nil {
		status = &models.Message{
			Id:       403,
			Name:     "Solicitações de Ignore-me devem ser confirmadas entre as 00:00 e as 23:59 do dia em que foram solicitadas!",
			Lang_key: "br",
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
	entity.Confirmed = true
	entity.Last_modified_date = time.Now().String()
	_, err := db.Update(entity)
	if err != nil {
		checkErr(err, "update failed")
	}
}

func InIgnoreList(db gorp.SqlExecutor, contact string) *models.Ignore_List {	

	ignored := models.Ignore_List{}

	err := db.SelectOne(&ignored, "SELECT * FROM ignore_list WHERE Contact= :contact ",  
		map[string]interface{}{
	  		"contact": contact, 	  		
		})
	checkErr(err, "select failed")
	

	return &ignored
}

func IsIgnoreRequestUnconfirmedSentToday(db gorp.SqlExecutor, contact string) bool {	

	str_today 	:= time.Now().Format(models.JsonFormat)

	select_stmt := "SELECT COUNT(*) FROM ignore_list "
	select_stmt += "WHERE Contact= ? AND confirmed = false "
	select_stmt += " Created_date BETWEEN '" + str_today + " 00:00:00' AND '" + str_today + " 23:59:59' "

	exists, err := db.SelectInt(select_stmt, contact)
	checkErr(err, "select failed")
	

	return exists > 0
}

//TODO: EXPIRAR A REQUISICAO DIARIAMENTE E REMOVER DO BANCO