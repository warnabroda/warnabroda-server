package models

//Captcha fields to validate over google's captcha server
type Captcha struct {
    Response string `json:"response"`   	//challenge resolved at view layer
	Ip string `json:"ip"`  					//challenge solver's IP
}
