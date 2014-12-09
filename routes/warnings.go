package routes

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/mostafah/mandrill"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	wab_root = os.Getenv("WARNAROOT")
	mandrill_key = os.Getenv("MANDRILL_KEY")
	mail_from = os.Getenv("WARNAEMAIL")
	wab_credencial = os.Getenv("WARNACREDENCIAL")
	wab_project = os.Getenv("WARNAPROJECT")
	sms_facilita_user = os.Getenv("WARNASMS_USER")
	sms_facilita_pass = os.Getenv("WARNASMS_PASS")
)

const (
	dbFormat   = "2006-01-02 15:04:05"
	jsonFormat = "2006-01-02"
)

//https://www.facilitamovel.com.br/api/simpleSend.ft?user=warnabroda&password=superwarnabroda951753&destinatario=4896662015&msg=WarnabrodaTest
func sendEmail(entity *models.Warning, db gorp.SqlExecutor) {

	mandrill.Key = mandrill_key
	// you can test your API key with Ping
	err := mandrill.Ping()
	// everything is OK if err is nil

	wab_email_template := wab_root + "/models/warning.html"

	//reads the e-mail template from a local file
	template_byte, err := ioutil.ReadFile(wab_email_template)
	checkErr(err, "File Opening ERROR")
	template_email_string := string(template_byte[:])

	//Get a random subject for the e-mails subject
	subject := GetRandomSubject()

	//Get the label for the sending warning
	message := SelectMessage(db, entity.Id_message)

	var email_content string
	email_content = strings.Replace(template_email_string, "{{warning}}", message.Name, 1)

	msg := mandrill.NewMessageTo(entity.Contact, subject.Name)
	msg.HTML = email_content
	// msg.Text = "plain text content" // optional
	msg.Subject = subject.Name
	msg.FromEmail = mail_from
	msg.FromName = "Warn A Broda: Dá um toque"

	//envio assincrono = true // envio sincrono = false
	res, err := msg.Send(false)
	fmt.Println(entity)
	if res[0] != nil {
		UpdateWarningSent(entity, db)
	} else {
		fmt.Println(res[0])
	}

}

//http://www.mpgateway.com/v_2_00/smspush/enviasms.aspx?CREDENCIAL=9A5A910B8B73AD14E5899D062B2A2AC2B065BD1B&PRINCIPAL_USER=WARNABRODA&AUX_USER=WAB&MOBILE=554896662015&SEND_PROJECT=N&MESSAGE=20Teste%20usando%20o%20mobile%20pronto2
func sendSMSMobilePronto(entity *models.Warning, db gorp.SqlExecutor){
	message := SelectMessage(db, entity.Id_message)
	sms_message := "Ola Amigo(a), "
	sms_message += "Você " + message.Name + ". "
	sms_message += "Avise um amigo você também: www.warnabroda.com"
	u, err := url.Parse("http://www.mpgateway.com/v_2_00/smsfollow/smsfollow.aspx?")

	if err != nil {
		checkErr(err, "Ugly URL")
	}
	u.Scheme = "http"
	u.Host = "www.mpgateway.com"
	q := u.Query()
	q.Set("CREDENCIAL", wab_credencial)
	q.Set("PRINCIPAL_USER", wab_project)
	q.Set("AUX_USER", "WAB")
	q.Set("MOBILE", "55"+entity.Contact)
	q.Set("SEND_PROJECT", "N")
	q.Set("MESSAGE", sms_message)
	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		checkErr(err, "SMS Not Sent")
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		checkErr(err, "No response from SMS Sender")
	} else {
		entity.Message = string(robots[:])
		UpdateWarningSent(entity, db)
	}
}
func sendSMS(entity *models.Warning, db gorp.SqlExecutor) {
	message := SelectMessage(db, entity.Id_message)
	sms_message := "Ola Amigo(a), "
	sms_message += "Você " + message.Name + ". "
	sms_message += "Avise um amigo você também: www.warnabroda.com"
	u, err := url.Parse("https://www.facilitamovel.com.br/api/simpleSend.ft?")

	if err != nil {
		checkErr(err, "Ugly URL")
	}
	u.Scheme = "https"
	u.Host = "www.facilitamovel.com.br"
	q := u.Query()
	q.Set("user", sms_facilita_user)
	q.Set("password", sms_facilita_pass)
	q.Set("destinatario", entity.Contact)
	q.Set("msg", sms_message)
	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		checkErr(err, "SMS Not Sent")
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		checkErr(err, "No response from SMS Sender")
	} else {
		entity.Message = string(robots[:])
		UpdateWarningSent(entity, db)
	}

}

func GetWarnings(enc Encoder, db gorp.SqlExecutor) (int, string) {
	var warnings []models.Warning
	_, err := db.Select(&warnings, "select * from warnings order by id")
	if err != nil {
		checkErr(err, "select failed")
		return http.StatusInternalServerError, ""
	}
	return http.StatusOK, Must(enc.Encode(warningsToIface(warnings)...))
}

func CountWarnings(enc Encoder, db gorp.SqlExecutor) (int, string) {
	
	total, err := db.SelectInt("SELECT COUNT(*) AS total FROM warnings WHERE Sent=true")
	if err != nil {
		checkErr(err, "COUNT ERROR")
		return http.StatusInternalServerError, ""
	}
	return http.StatusOK, strconv.FormatInt(total, 10)
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
	if err != nil {
		checkErr(err, "insert failed")
		return http.StatusConflict, ""
	}
	w.Header().Set("Location", fmt.Sprintf("/warnabroda/warnings/%d", entity.Id))

	if entity.Id_contact_type == 1 {
		processEmail(&entity, db, status)
	} else if entity.Id_contact_type == 2 {
		processSMS(&entity, db, status)

	}

	return http.StatusCreated, Must(enc.EncodeOne(status))
}

func processEmail(warning *models.Warning, db gorp.SqlExecutor, status *models.Message){
	fmt.Println(warning)
	if emailSentToContact(warning, db) {
		status.Id = 403
		status.Name = "Broda já foi avisado(a) há instantes atrás. Muito Obrigado."
		status.Lang_key = "br"
	} else {
		go sendEmail(warning, db)
	}
}

func processSMS(warning *models.Warning, db gorp.SqlExecutor, status *models.Message) {

	if smsSentToContact(warning, db) {
		status.Id = 403
		status.Name = "Este número já recebeu um SMS hoje ou seu IP(" + warning.Ip + ") já enviou a cota maxima de SMS diário."
		status.Lang_key = "br"
	} else {
		go sendSMSMobilePronto(warning, db)
	}

}

func emailSentToContact(warning *models.Warning, db gorp.SqlExecutor) bool {
	
	now_lower := time.Now().Add(-2*time.Hour)
	now_upper := time.Now().Add(2*time.Hour)

	str_now_lower			:= now_lower.Format(dbFormat)			
	str_now_upper			:= now_upper.Format(dbFormat)	
	
	return_statement 	:= false
	var warnings []models.Warning
	
	select_statement := " SELECT * FROM warnings "
	select_statement += " WHERE Id_contact_type = 1 AND Sent = 1 AND "
	select_statement += " Contact = '" + warning.Contact + "' AND "
	select_statement += " Created_date BETWEEN '" + str_now_lower +"' AND '" + str_now_upper +"' AND "
	select_statement += " Id <> " + strconv.FormatInt(warning.Id, 10)
	fmt.Println(select_statement)

	_, err := db.Select(&warnings, select_statement)
	if err != nil {
		checkErr(err, "Checking Contact failed")
	}

	if len(warnings) > 0 {
		return_statement = true
	}

	return return_statement
}



func smsSentToContact(warning *models.Warning, db gorp.SqlExecutor) bool {

	str_today := time.Now().Format(jsonFormat)

	return_statement := false
	var warnings []models.Warning

	select_statement := " SELECT * FROM warnings "
	select_statement += " WHERE Id_contact_type = 2 AND Sent = 1 AND "
	select_statement += " (Contact = '" + warning.Contact + "' OR Ip LIKE '%" + warning.Ip + "%' ) AND "
	select_statement += " Created_date BETWEEN '" + str_today + " 00:00:00' AND '" + str_today + " 23:59:59' AND "
	select_statement += " Id <> " + strconv.FormatInt(warning.Id, 10)

	_, err := db.Select(&warnings, select_statement)
	if err != nil {
		checkErr(err, "Checking Contact failed")
	}

	if len(warnings) > 0 {
		return_statement = true
	}

	return return_statement
}

func UpdateWarningSent(entity *models.Warning, db gorp.SqlExecutor) {
	entity.Sent = true
	entity.Last_modified_date = time.Now().String()
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
