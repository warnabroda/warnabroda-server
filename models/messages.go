package models



type Messages struct {
    Id 			int64  `json:"id" 		db:"id"`    			// Id cast as int64
    Name 		string `json:"name"		db:"name"`  			// Name cast as string
    Lang_key 	string `json:"lang_key"	form:"lang_key" db:"lang_key"`  	// Lang_key language used at Name field
	Total 		string `json:"total"	db:"total"`
	NotSent 	string `json:"not_sent"	form:"not_sent" db:"not_sent"`
	Sent 		string `json:"sent"		db:"sent"`
	Email 		string `json:"email"	db:"email"`
	Sms 		string `json:"sms"		db:"sms"`
	Whatsapp 	string `json:"whatsapp"	db:"whatsapp"`
}


