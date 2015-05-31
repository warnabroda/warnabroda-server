package routes

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	// "fmt"
	"bytes"
	"encoding/json"

	"github.com/coopernurse/gorp"
	"gitlab.com/warnabroda/warnabrodagomartini/models"
)

const (
	GOOGLE_SHORTNER_URL = "https://www.googleapis.com/urlshortener/v1/url?"
	GOOGLE_CAPTCHA_URL  = "https://www.google.com/recaptcha/api/siteverify?"
	SCHEME              = "https"
	HOST                = "www.google.com"
	API_HOST            = "www.googleapis.com"
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

	status := &models.DefaultStruct{
		Id:       http.StatusBadRequest,
		Name:     "Warning Update Failed.",
		Lang_key: "en",
	}

	switch entity.Type {
	case "warning":

		warning := GetWarning(entity.Id, db)
		if warning != nil {
			warning.Message = entity.Name
			if UpdateWarningSent(warning, db) {
				status.Id = http.StatusAccepted
				status.Name = "Warning Update Success."
				status.Lang_key = warning.Lang_key
			}
		}
	case "ignore":

		ignoreItem := GetIgnoreContactById(entity.Id, db)
		if ignoreItem != nil {
			if UpdateIgnoreSent(ignoreItem, db) {
				status.Id = http.StatusAccepted
				status.Name = "IgnoreList Update Success."
				status.Lang_key = ignoreItem.Lang_key
			}
		}

	case "reply":
		reply := GetReplyById(entity.Id, db)
		if reply != nil {
			if UpdateReplySent(reply, db) {
				status.Id = http.StatusAccepted
				status.Name = "Reply Update Success."
				status.Lang_key = reply.Lang_key
			}
		}

	}

	return http.StatusOK, Must(enc.EncodeOne(status))
}

func ShortUrl(longLink string) string {
	u, err := url.Parse(GOOGLE_SHORTNER_URL)
	checkErr(err, "Ugly URL")

	u.Scheme = SCHEME
	u.Host = API_HOST
	q := u.Query()

	q.Set("key", "AIzaSyB_ZX0ebsr5RxV8UvfPZlC6Obp3i3dqemw")

	u.RawQuery = q.Encode()

	res, err := http.Post(u.String(), "application/json", bytes.NewReader([]byte(`{"longUrl":"`+longLink+`"}`)))

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	checkErr(err, "No response from Google Captcha Server")

	var data models.GoogleShortner
	err = json.Unmarshal(body, &data)

	return data.Id
}
