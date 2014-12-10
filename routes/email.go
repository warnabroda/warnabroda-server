package routes

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"github.com/coopernurse/gorp"
	"github.com/mostafah/mandrill"
	"io/ioutil"
	"os"
)

var (	
	mandrill_key 		= os.Getenv("MANDRILL_KEY")
	mail_from 			= os.Getenv("WARNAEMAIL")
)

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
