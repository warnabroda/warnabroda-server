package routes

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"bitbucket.org/hbtsmith/warnabrodagomartini/i18n"
	"github.com/coopernurse/gorp"
	"github.com/mostafah/mandrill"
	"io/ioutil"
	"os"
	"strings"
	"fmt"
)

var (	
	mandrill_key 		= os.Getenv("MANDRILL_KEY")
	mail_from 			= os.Getenv("WARNAEMAIL")
	wab_root 			= os.Getenv("WARNAROOT")
)


func ProcessEmail(warning *models.Warning, db gorp.SqlExecutor){	
	
	go sendEmailWarn(warning, db)
	
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
		fmt.Println("SENDING MAIL (sendEmailWarn) ERROR")
	}

}

func SendMail(email *models.Email, db gorp.SqlExecutor) bool {

	mandrill.Key = mandrill_key
	// you can test your API key with Ping
	err := mandrill.Ping()
	// everything is OK if err is nil
	msg := mandrill.NewMessageTo(email.ToAddress, email.Subject)

	var content string

	if email.UseContent {
		content = email.Content
	} else {
		//reads the e-mail template from a local file
		template_byte, err := ioutil.ReadFile(email.TemplatePath)
		checkErr(err, "File Opening ERROR")
		content = string(template_byte[:])
	}

	if (email.HTMLContent){
		msg.HTML = content
	} else {
		msg.Text = content
	}	
	
	msg.Subject = email.Subject
	msg.FromEmail = mail_from
	msg.FromName = email.FromName

	//envio assincrono = true // envio sincrono = false
	res, err := msg.Send(email.Async)
	checkErr(err, "File Opening ERROR")		
	
	return res[0] != nil
}
