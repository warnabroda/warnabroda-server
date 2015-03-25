package routes

import (
	"io/ioutil"
	"encoding/json"
	"os"
	"strings"
	"fmt"

	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"bitbucket.org/hbtsmith/warnabrodagomartini/messages"
	"github.com/coopernurse/gorp"
	"github.com/mostafah/mandrill"
)

var (	
	mandrill_key 		= os.Getenv("MANDRILL_KEY")
	mail_from 			= os.Getenv("WARNAEMAIL")
	wab_root 			= os.Getenv("WARNAROOT")
)


// For now all due verifications regarding send rules is done previewsly, here we just async the e-mail send of the warn
func ProcessEmail(entity *models.Warning, db gorp.SqlExecutor){
	
	go sendEmailWarn(entity, db)
}

func SendEmailReplyRequestAcknowledge(entity *models.WarningResp, db gorp.SqlExecutor){
	
	//reads the e-mail template from a local file
	wab_email_template := wab_root + "/resource/reply_ack_"+entity.Lang_key+".html"
	template_byte, err := ioutil.ReadFile(wab_email_template)
	checkErr(err, "Email File Opening ERROR")
	template_email_string := string(template_byte[:])

	subject  := messages.GetLocaleMessage(entity.Lang_key, "MSG_SUBJECT_REPLY_REQUEST")
	mail_msg := strings.Replace(messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_SENDER_MSG"), "{{reply_to}}", entity.Reply_to, 1)
	
	var email_content string
	email_content = strings.Replace(template_email_string, "{{reply_ack_msg}}", mail_msg, 1)
	email_content = strings.Replace(email_content, "{{url_contacus}}", models.URL_CONTACT_US, 1)
	email_content = strings.Replace(email_content, "{{email}}", models.EMAIL_WARNABRODA, 2)

	email := &models.Email{
		TemplatePath: wab_email_template,	
		Content: email_content, 	
		Subject: subject,		
		ToAddress: entity.Reply_to,
		FromName: models.WARN_A_BRODA,
		LangKey: entity.Lang_key,
		Async: false,
		UseContent: true,
		HTMLContent: true,
	}	
	
	sent, _ := SendMail(email)

	if !sent {
		fmt.Println("SENDING MAIL (sendEmailWarn) ERROR")
	}
	
	

}

//Deploys the message to be sent into an email struct, call the service and in case of successful send, update the warn as sent.
func sendEmailWarn(entity *models.Warning, db gorp.SqlExecutor) {	

	//reads the e-mail template from a local file
	wab_email_template := wab_root + "/resource/warning_"+entity.Lang_key+".html"
	template_byte, err := ioutil.ReadFile(wab_email_template)
	checkErr(err, "Email File Opening ERROR")
	template_email_string := string(template_byte[:])

	subject := GetRandomSubject(entity.Lang_key)
	message := SelectMessage(db, entity.Id_message)
	var email_content string
	email_content = strings.Replace(template_email_string, "{{warning}}", message.Name, 1)
	email_content = strings.Replace(email_content, "{{url_contacus}}", models.URL_CONTACT_US, 1)
	email_content = strings.Replace(email_content, "{{email}}", models.EMAIL_WARNABRODA, 2)

	if entity.WarnResp != nil {
		if entity.WarnResp.Id_contact_type == 1 {
			go SendEmailReplyRequestAcknowledge(entity.WarnResp, db)
		} else {
			SendWhatsappReplyRequestAcknowledge(entity.WarnResp, db)
		}

		
		email_content = strings.Replace(email_content, "{{reply_to}}", models.URL_REPLY_TO + "/" + entity.WarnResp.Resp_hash, 2)

	}

	email := &models.Email{
		TemplatePath: wab_email_template,	
		Content: email_content, 	
		Subject: subject.Name,		
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
		UpdateWarningSent(entity, db)
	} else {
		fmt.Println("SENDING MAIL (sendEmailWarn) ERROR")
	}



}

func SendMail(email *models.Email) (bool, string) {

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
	checkErr(err, "SendMail File Opening ERROR")
	resp, _ := json.Marshal(res[0])	
	
	return res[0] != nil, string(resp)
}
