package routes

import (
	"bytes"
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	//	"time"

	"github.com/coopernurse/gorp"
	"gitlab.com/warnabroda/warnabrodagomartini/models"
	"gopkg.in/redis.v3"
)

var (
	GoogleShortUrlKey = os.Getenv("WARNA_GOOGLE_SHORTURL_KEY")
	client            = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
)

const (
	GOOGLE_SHORTNER_URL = "https://www.googleapis.com/urlshortener/v1/url?"
	GOOGLE_CAPTCHA_URL  = "https://www.google.com/recaptcha/api/siteverify?"
	SCHEME              = "https"
	HOST                = "www.google.com"
	API_HOST            = "www.googleapis.com"
	DISTINCT_PHONE_SQL  = " SELECT distinct wr.reply_to AS contact FROM warning_resp AS wr WHERE wr.id_contact_type <> 1 "
)

func init() {
	LoadPhoneRegis()
}

func LoadPhoneRegis() {
	var phones []string
	_, err := models.Dbm.Select(&phones, DISTINCT_PHONE_SQL)

	_, erredis := client.Ping().Result()

	if err == nil && erredis == nil {
		i := 0
		for i < len(phones) {
			//fmt.Println()
			err := client.Set(phones[i], "true", 0).Err()
			if err != nil {
				checkErr(err, "Fail loading Redis data")
			}
			i++
		}
	}

	//	val, err := client.Get("+55489666201554").Result()
	//	if err != nil {
	//		fmt.Println("É dá erro mesmo")
	//		fmt.Println(val)
	//		fmt.Println(err)
	//	}
	//	fmt.Println("+55489666201554", val)

}

func AddPhoneToRedis(phone string) {
	err := client.Set(phone, "true", 0).Err()
	if err != nil {
		checkErr(err, "Fail loading Redis data")
	}
}

func IsLoadedInRedis(phone string) bool {
	val, err := client.Get(phone).Result()
	return val != "" && err == nil
}

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

	q.Set("key", GoogleShortUrlKey)

	u.RawQuery = q.Encode()

	res, err := http.Post(u.String(), "application/json", bytes.NewReader([]byte(`{"longUrl":"`+longLink+`"}`)))

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	checkErr(err, "No response from Google Captcha Server")

	var data models.GoogleShortner
	err = json.Unmarshal(body, &data)

	return data.Id
}
