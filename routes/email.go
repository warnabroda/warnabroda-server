package routes

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	// "fmt"

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

	go sendEmailWarn(entity, db)
}

func SendEmailReplyDone(entity *models.Warning, db gorp.SqlExecutor) {

	//reads the e-mail template from a local file
	wab_email_template := wab_root + "/resource/reply_" + entity.Lang_key + ".html"
	template_byte, err := ioutil.ReadFile(wab_email_template)
	checkErr(err, "Email File Opening ERROR")
	template_email_string := string(template_byte[:])
	message := SelectMessage(db, entity.Id_message)

	subject := messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_EMAIL_SUBJECT") + ": " + entity.WarnResp.Resp_hash[:6]

	mainBodyMsg := messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_BODY_MAIN")
	mainBodyMsg = strings.Replace(mainBodyMsg, "{{sent_msg}}", "'"+message.Name+"'", 1)
	mainBodyMsg = strings.Replace(mainBodyMsg, "{{contact_msg}}", "'"+entity.Contact+"'", 1)

	var email_content string
	email_content = strings.Replace(template_email_string, "{{greeting}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_SMS_HEADER"), 1)
	email_content = strings.Replace(email_content, "{{body_header}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_BODY_GREETING"), 1)
	email_content = strings.Replace(email_content, "{{body_main_content}}", mainBodyMsg, 1)
	email_content = strings.Replace(email_content, "{{body_footer_content}}", messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_BODY_LINK"), 1)
	email_content = strings.Replace(email_content, "{{reply_url}}", models.URL_REPLY+"/"+entity.WarnResp.Read_hash, 2)

	email_content = strings.Replace(email_content, "{{url_contacus}}", models.URL_CONTACT_US, 1)
	email_content = strings.Replace(email_content, "{{email}}", models.EMAIL_WARNABRODA, 2)

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
	wab_email_template := wab_root + "/resource/reply_ack_" + entity.Lang_key + ".html"
	template_byte, err := ioutil.ReadFile(wab_email_template)
	checkErr(err, "Email File Opening ERROR")
	template_email_string := string(template_byte[:])

	subject := messages.GetLocaleMessage(entity.Lang_key, "MSG_SUBJECT_REPLY_REQUEST") + ": " + entity.Resp_hash[:6]
	mail_msg := strings.Replace(messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_SENDER_MSG"), "{{reply_to}}", entity.Reply_to, 1)

	var email_content string
	email_content = strings.Replace(template_email_string, "{{reply_ack_msg}}", mail_msg, 1)
	email_content = strings.Replace(email_content, "{{url_contacus}}", models.URL_CONTACT_US, 1)
	email_content = strings.Replace(email_content, "{{email}}", models.EMAIL_WARNABRODA, 2)

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
	wab_email_template := wab_root + "/resource/warning_" + entity.Lang_key + ".html"
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
			go SendWhatsappReplyRequestAcknowledge(entity.WarnResp, db)
		}

		replyLink := messages.GetLocaleMessage(entity.Lang_key, "MSG_FOOTER_REPLY")
		replyToMsgLink := strings.Replace(replyLink, "{{url_reply}}", models.URL_REPLY+"/"+entity.WarnResp.Resp_hash, 1)
		replyToContent := "<h4><a target='_blank' href='" + models.URL_REPLY + "/" + entity.WarnResp.Resp_hash + "'>" + replyToMsgLink + "</a></h4><br/>"
		email_content = strings.Replace(email_content, "{{reply_to}}", replyToContent, 1)

	} else {
		email_content = strings.Replace(email_content, "{{reply_to}}", "", 1)
	}

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
