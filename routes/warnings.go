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
	"strconv"
	"strings"
	"time"
	"os"
)

//https://www.facilitamovel.com.br/api/simpleSend.ft?user=warnabroda&password=superwarnabroda951753&destinatario=4896662015&msg=WarnabrodaTest
func sendEmail(entity *models.Warning, db gorp.SqlExecutor) {	

	mandrill.Key = "qUX983QXREtaLojEpJyxmw"
	// you can test your API key with Ping
	err := mandrill.Ping()
	// everything is OK if err is nil
	
	var wab_root = os.Getenv("WARNAROOT")
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
	msg.FromEmail = "warnabroda@gmail.com"
	msg.FromName = "Warn A Broda: Dá um toque"

	//envio assincrono = true // envio sincrono = false
	res, err := msg.Send(false)

	if res[0] != nil {
		UpdateWarningSent(entity, db)
	} else {
		fmt.Println(res[0])
	}
	
}

func sendSMS(entity *models.Warning, db gorp.SqlExecutor){
	message := SelectMessage(db, entity.Id_message)	
	sms_message := "Ola Amigo(a), "
	sms_message += "Você " + message.Name + ". "
	sms_message += "Avise um amigo você também: www.warnabroda.com"
	u, err := url.Parse("https://www.facilitamovel.com.br/api/simpleSend.ft?")
	fmt.Println(u)
	if err != nil {
		checkErr(err, "Ugly URL")
	}
	u.Scheme = "https"
	u.Host = "www.facilitamovel.com.br"
	q := u.Query()
	q.Set("user", "user")//warnabroda
	q.Set("password", "password")//superwarnabroda951753
	q.Set("destinatario", entity.Contact)
	q.Set("msg",  sms_message)
	u.RawQuery = q.Encode()
	fmt.Println(u.String())

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
    	fmt.Printf("%s", robots)
    	fmt.Println("#####")
    	fmt.Println(entity)
    	fmt.Println("#####")
    }


	//?user=warnabroda&password=superwarnabroda951753&destinatario=4896662015&msg=WarnabrodaTest
	//https://www.facilitamovel.com.br/api/simpleSend.ft?destinatario=4896662015&msg=Amigo%28a%29%2C+Voc%C3%AA+tem+sujeira+de+merda+grudada+no+sanit%C3%A1rio.+Avise+um+amigo+voc%C3%AA+tam%C3%A9m%3A+www.warnabroda.com&password=superwarnabroda951753&user=warnabroda

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

	entity.Sent = false
	entity.Created_by = "system"
	entity.Created_date = models.JDate(time.Now())
	entity.Lang_key = "br"

	err := db.Insert(&entity)
	if err != nil {
		checkErr(err, "insert failed")
		return http.StatusConflict, ""
	}
	w.Header().Set("Location", fmt.Sprintf("/warnabroda/warnings/%d", entity.Id))
	if entity.Id_contact_type == 1 {
		go sendEmail(&entity, db)
	} else if entity.Id_contact_type == 2 {
		go sendSMS(&entity, db)
	}
	return http.StatusCreated, Must(enc.EncodeOne(entity))
}

func UpdateWarningSent(entity *models.Warning, db gorp.SqlExecutor) {
	entity.Sent = true
	entity.Last_modified_date = models.JDate(time.Now())
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
