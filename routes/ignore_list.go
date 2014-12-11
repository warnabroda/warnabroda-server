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

	if InIgnoreList(db, entity.Contact){
		status = &models.Message{
			Id:       403,
			Name:     "Contato Já estava na Lista de Ignorados!",
			Lang_key: "br",
		}		
	} else {
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

		err := db.Insert(&entity)
		if err != nil {		
			checkErr(err, "INSERT IGNORE FAIL")
		}
	}

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

func InIgnoreList(db gorp.SqlExecutor, contact string) bool {

	exists, err := db.SelectInt("SELECT COUNT(*) FROM ignore_list WHERE Contact=? AND confirmed=true", contact)

	if err != nil {
		checkErr(err, "select failed")
	}

	return exists > 0
}