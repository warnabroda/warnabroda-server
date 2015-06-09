package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/coopernurse/gorp"
	"github.com/mostafah/mandrill"
	"gitlab.com/warnabroda/warnabrodagomartini/messages"
	"gitlab.com/warnabroda/warnabrodagomartini/models"
)

var (
	mandrill_key = os.Getenv("MANDRILL_KEY")
	mail_from    = os.Getenv("WARNAEMAIL")
	wab_root     = os.Getenv("WARNAROOT")
)

// For now all due verifications regarding send rules is done previewsly, here we just async the e-mail send of the warn
func ProcessEmail(entity *models.Warning, db gorp.SqlExecutor) {
	sendEmailWarn(entity, db)
}

func SendEmailReplyDone(entity *models.Warning, db gorp.SqlExecutor) {

	//reads the e-mail template from a local file
	wab_email_template := wab_root + "/resource/reply.html"
	template_byte, err := ioutil.ReadFile(wab_email_template)
	checkErr(err, "Email File Opening ERROR")
	template_email_string := string(template_byte[:])

	subject := messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_EMAIL_SUBJECT") + ": " + entity.WarnResp.Resp_hash[:6]

	email_content := sendReplySetup(template_email_string, entity, db)

	email := &models.Email{
		TemplatePath: wab_email_template,
		Content:      email_content,
		Subject:      subject,
		ToAddress:    entity.WarnResp.Reply_to,
		FromName:     models.WARN_A_BRODA,
		LangKey:      entity.Lang_key,
		Async:        false,
		UseContent:   true,
		HTMLContent:  true,
	}

	sent, _ := SendMail(email)

	if sent {
		UpdateReplySent(entity.WarnResp, db)
	}
}

func SendEmailReplyRequestAcknowledge(entity *models.WarningResp, db gorp.SqlExecutor) {

	//reads the e-mail template from a local file
	wab_email_template := wab_root + "/resource/warning.html"
	template_byte, err := ioutil.ReadFile(wab_email_template)
	checkErr(err, "Email File Opening ERROR")
	template_email_string := string(template_byte[:])

	subject := messages.GetLocaleMessage(entity.Lang_key, "MSG_SUBJECT_REPLY_REQUEST") + ": " + entity.Resp_hash[:6]

	email_content := sendReplyAcknowledgeSetup(template_email_string, entity, db)

	email := &models.Email{
		TemplatePath: wab_email_template,
		Content:      email_content,
		Subject:      subject,
		ToAddress:    entity.Reply_to,
		FromName:     models.WARN_A_BRODA,
		LangKey:      entity.Lang_key,
		Async:        false,
		UseContent:   true,
		HTMLContent:  true,
	}

	SendMail(email)
}

//Deploys the message to be sent into an email struct, call the service and in case of successful send, update the warn as sent.
func sendEmailWarn(entity *models.Warning, db gorp.SqlExecutor) {

	//reads the e-mail template from a local file
	wab_email_template := wab_root + "/resource/warning.html"

	if entity.WarnResp != nil {
		wab_email_template = wab_root + "/resource/warning-reply.html"
	}
	template_byte, err := ioutil.ReadFile(wab_email_template)
	checkErr(err, "Email File Opening ERROR")
	template_email_string := string(template_byte[:])

	subject := GetRandomSubject(entity.Lang_key)

	email_content := sendWarningSetup(template_email_string, entity, db)

	email := &models.Email{
		TemplatePath: wab_email_template,
		Content:      email_content,
		Subject:      subject.Name,
		ToAddress:    entity.Contact,
		FromName:     models.WARN_A_BRODA,
		LangKey:      entity.Lang_key,
		Async:        false,
		UseContent:   true,
		HTMLContent:  true,
	}

	SendMail(email)
	sent, response := SendMail(email)

	if sent {
		entity.Message = response
		UpdateWarningSent(entity, db)
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

	if email.HTMLContent {
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

func sendWarningSetup(email string, entity *models.Warning, db gorp.SqlExecutor) string {

	var email_content string
	subject := GetRandomSubject(entity.Lang_key)
	message := SelectMessage(db, entity.Id_message)
	email_content = strings.Replace(email, "{{subject}}", subject.Name, 1)
	email_content = strings.Replace(email_content, "{{broda_msg_header_greet}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_EMAIL_GREET"), 1)
	email_content = strings.Replace(email_content, "{{warning}}", message.Name, 1)
	email_content = strings.Replace(email_content, "{{broda_msg_footer_greet}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_FOOTER"), 1)
	email_content = strings.Replace(email_content, "{{warnabroda_headline}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_WARNABRODA_HEADLINE"), 1)
	email_content = strings.Replace(email_content, "{{follow_us}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_FOLLOW_US"), 1)
	email_content = strings.Replace(email_content, "{{terms}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_TERMS_SERVICE"), 1)

	if entity.WarnResp != nil {
		if entity.WarnResp.Id_contact_type == 1 {
			SendEmailReplyRequestAcknowledge(entity.WarnResp, db)
		} else {
			SendWhatsappReplyRequestAcknowledge(entity, db)
		}

		email_content = strings.Replace(email_content, "{{reply_url}}", ShortUrl(models.URL_REPLY+"/"+entity.WarnResp.Resp_hash), 1)
		email_content = strings.Replace(email_content, "{{reply_msg}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_NOW"), 1)
	}

	the7messages := GetRandomMessagesByLanguage(7, entity.Lang_key, db)

	for key, value := range the7messages {

		email_content = strings.Replace(email_content, "{{suggest_msg"+fmt.Sprint(key+1)+"}}", value.Name, 1)
		email_content = strings.Replace(email_content, "{{suggest_link"+fmt.Sprint(key+1)+"}}", ShortUrl(models.URL_WARNABRODA+"/#/"+fmt.Sprint(value.Id)), 1)
	}

	return email_content

}

func sendReplyAcknowledgeSetup(email string, entity *models.WarningResp, db gorp.SqlExecutor) string {

	var email_content string
	email_content = strings.Replace(email, "{{subject}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_SUBJECT_REPLY_REQUEST"), 1)
	email_content = strings.Replace(email_content, "{{broda_msg_header_greet}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_SENDER_MSG"), 1)
	email_content = strings.Replace(email_content, "{{warning}}", entity.Reply_to, 1)
	email_content = strings.Replace(email_content, "{{broda_msg_footer_greet}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_FOOTER"), 1)
	email_content = strings.Replace(email_content, "{{warnabroda_headline}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_HEADLINE"), 1)
	email_content = strings.Replace(email_content, "{{follow_us}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_FOLLOW_US"), 1)
	email_content = strings.Replace(email_content, "{{terms}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_TERMS_SERVICE"), 1)

	the7messages := GetRandomMessagesByLanguage(7, entity.Lang_key, db)

	for key, value := range the7messages {

		email_content = strings.Replace(email_content, "{{suggest_msg"+fmt.Sprint(key+1)+"}}", value.Name, 1)
		email_content = strings.Replace(email_content, "{{suggest_link"+fmt.Sprint(key+1)+"}}", ShortUrl(models.URL_WARNABRODA+"/#/"+fmt.Sprint(value.Id)), 1)
	}

	return email_content

}

func sendReplySetup(email string, entity *models.Warning, db gorp.SqlExecutor) string {

	var email_content string
	email_content = strings.Replace(email, "{{subject}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_BODY_GREETING"), 1)

	message := SelectMessage(db, entity.Id_message)

	mainBodyMsg := messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_BODY_MAIN")
	mainBodyMsg = strings.Replace(mainBodyMsg, "{{contact_msg}}", "'"+entity.Contact+"'", 1)

	email_content = strings.Replace(email_content, "{{broda_msg_header_greet}}", mainBodyMsg, 1)
	email_content = strings.Replace(email_content, "{{warning}}", message.Name, 1)
	email_content = strings.Replace(email_content, "{{broda_msg_footer_greet}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_READ_MSG"), 1)
	email_content = strings.Replace(email_content, "{{warnabroda_headline}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_WARNABRODA_HEADLINE"), 1)
	email_content = strings.Replace(email_content, "{{follow_us}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_FOLLOW_US"), 1)
	email_content = strings.Replace(email_content, "{{terms}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_TERMS_SERVICE"), 1)

	email_content = strings.Replace(email_content, "{{reply_url}}", ShortUrl(models.URL_REPLY+"/"+entity.WarnResp.Read_hash), 1)
	email_content = strings.Replace(email_content, "{{reply_msg}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_READ_REPLY_NOW"), 1)

	the7messages := GetRandomMessagesByLanguage(7, entity.Lang_key, db)

	for key, value := range the7messages {

		email_content = strings.Replace(email_content, "{{suggest_msg"+fmt.Sprint(key+1)+"}}", value.Name, 1)
		email_content = strings.Replace(email_content, "{{suggest_link"+fmt.Sprint(key+1)+"}}", ShortUrl(models.URL_WARNABRODA+"/#/"+fmt.Sprint(value.Id)), 1)
	}

	return email_content

}

// opens the template, parse the variables sets the email struct and Send the confirmation code to confirm the ignored contact.
func SendEmailIgnoreme(entity *models.Ignore_List, db gorp.SqlExecutor) {
	//reads the e-mail template from a local file
	wab_email_template := wab_root + "/resource/ignoreme.html"
	template_byte, err := ioutil.ReadFile(wab_email_template)
	checkErr(err, "Ignore-me Email File Opening ERROR")
	template_email_string := string(template_byte[:])

	email_content := sendIgnoreMeSetup(template_email_string, entity, db)

	email := &models.Email{
		TemplatePath: wab_email_template,
		Content:      email_content,
		Subject:      messages.GetLocaleMessage(entity.Lang_key, "MSG_EMAIL_SUBJECT_ADD_IGNORE_LIST"),
		ToAddress:    entity.Contact,
		FromName:     models.WARN_A_BRODA,
		LangKey:      entity.Lang_key,
		Async:        false,
		UseContent:   true,
		HTMLContent:  true,
	}

	sent, response := SendMail(email)
	if sent {
		entity.Message = response
		UpdateIgnoreList(entity, db)
	}
}

func sendIgnoreMeSetup(email string, entity *models.Ignore_List, db gorp.SqlExecutor) string {

	var email_content string
	email_content = strings.Replace(email, "{{subject}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_IGNOREME_SUBJECT"), 1)

	email_content = strings.Replace(email_content, "{{broda_msg_header_greet}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_IGNOREME_GREET"), 1)
	email_content = strings.Replace(email_content, "{{warning}}", entity.Confirmation_code, 1)
	email_content = strings.Replace(email_content, "{{broda_msg_footer_greet}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_IGNOREME_FOOTER"), 1)
	email_content = strings.Replace(email_content, "{{warnabroda_headline}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_WARNABRODA_HEADLINE"), 1)
	email_content = strings.Replace(email_content, "{{follow_us}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_FOLLOW_US"), 1)
	email_content = strings.Replace(email_content, "{{terms}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_TERMS_SERVICE"), 1)

	// email_content = strings.Replace(email_content, "{{reply_url}}", ShortUrl(models.URL_IGNOREME+"/"+entity.WarnResp.Read_hash), 1)
	email_content = strings.Replace(email_content, "{{reply_url}}", ShortUrl(models.URL_IGNOREME), 1)
	email_content = strings.Replace(email_content, "{{reply_msg}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_IGNOREME_BUTTON"), 1)
	email_content = strings.Replace(email_content, "{{broda_msg_ignoreme_footer}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_IGNOREME_NOT"), 1)

	the7messages := GetRandomMessagesByLanguage(7, entity.Lang_key, db)

	for key, value := range the7messages {

		email_content = strings.Replace(email_content, "{{suggest_msg"+fmt.Sprint(key+1)+"}}", value.Name, 1)
		email_content = strings.Replace(email_content, "{{suggest_link"+fmt.Sprint(key+1)+"}}", ShortUrl(models.URL_WARNABRODA+"/#/"+fmt.Sprint(value.Id)), 1)
	}

	return email_content

}
