package models

type GoogleShortner struct {
    Kind string `json:"kind"`   	//    "id": "http://goo.gl/d81AfD",
	Id string `json:"id"`			//    "kind": "urlshortener#url",
	LongUrl string `json:"longUrl"` //    "longUrl": "http://www.warnabroda.com/"
}

//Captcha fields to validate over google's captcha server
type Captcha struct {
    Response string `json:"response"`   	//challenge resolved at view layer
	Ip string `json:"ip"`  					//challenge solver's IP
}