package routes

import (
	"net/http"
	"io/ioutil"
	"net/url"
	"os"

	"fmt"

	"bitbucket.org/hbtsmith/warnabrodagomartini/models"	
	"github.com/coopernurse/gorp"
)

const(
	GOOGLE_CAPTCHA_URL		= "https://www.google.com/recaptcha/api/siteverify?"
	SCHEME 					= "https"
	HOST 					= "www.google.com"
)

// Get google's captcha response
func CaptchaResponse(captcha models.Captcha, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {
	u, err := url.Parse(GOOGLE_CAPTCHA_URL)
	checkErr(err, "Ugly URL")	

	u.Scheme = SCHEME
	u.Host = HOST
	q := u.Query()

	q.Set("secret", os.Getenv("WARNACAPTCHA"))
	q.Set("response", captcha.Response)
	q.Set("remoteip", captcha.Ip)			
	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())	
	checkErr(err, "Captcha not verified")
	
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	checkErr(err, "No response from Google Captcha Server")
	
	
	return http.StatusOK, string(robots[:])
}

func SendConfirmation(entity models.DefaultStruct, enc Encoder, db gorp.SqlExecutor) (int, string) {	
		
	status 			:= &models.DefaultStruct{
		Id:       	http.StatusNotFound,
		Name:     	"Warning Update Failed.",
		Lang_key: 	"en",
	}

	fmt.Println(entity.Type)

	switch entity.Type {
		case "warning":

			warning 		:= GetWarning(entity.Id, db)
			if warning != nil{
				warning.Message 	= entity.Name
				if UpdateWarningSent(warning, db) {
					status.Id 			= http.StatusAccepted
					status.Name 		= "Warning Update Success."
					status.Lang_key 	= warning.Lang_key
				}
			}
		case "ignore":
			
		case "reply":
			
	}

	return http.StatusOK, Must(enc.EncodeOne(status))
}



