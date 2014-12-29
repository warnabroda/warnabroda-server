package routes

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"github.com/coopernurse/gorp"
//	"io/ioutil"
//	"net/http"
	"net/url"
	"fmt"
)

// Component to send a SMS using mobile pronto
func SendSMS(sms *models.SMS, db gorp.SqlExecutor) (bool, string) {
	u, err := url.Parse(sms.URLPath)
	checkErr(err, "Ugly URL")

	u.Scheme = sms.Scheme
	u.Host = sms.Host
	q := u.Query()
	q.Set("CREDENCIAL", sms.CredencialKey)
	q.Set("PRINCIPAL_USER", sms.Project)
	q.Set("AUX_USER", sms.AuxUser)
	q.Set("MOBILE", sms.MobileNumber)
	q.Set("SEND_PROJECT", sms.SendProject)
	q.Set("MESSAGE", sms.Content)
	u.RawQuery = q.Encode()

	fmt.Println(u)

	//res, err := http.Get(u.String())	
	// checkErr(err, "SMS Not Sent")

	//robots, err := ioutil.ReadAll(res.Body)
	//res.Body.Close()
	// checkErr(err, "No response from SMS Sender")

	return true, "TESTE"
	//return err == nil, string(robots[:])
}