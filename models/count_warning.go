package models

type CountWarning struct {
	Allw 				string `json:"All" 					db:"Allw"`
	Sent 				string `json:"Sent" 				db:"Sent"`
	NotSent				string `json:"NotSent" 				db:"NotSent"`
	SmsSent				string `json:"SmsSent" 				db:"SmsSent"`	
	SmsNotSent			string `json:"SmsNotSent" 			db:"SmsNotSent"`
	EmailSent			string `json:"EmailSent" 			db:"EmailSent"`
	EmailNotSent		string `json:"EmailNotSent" 		db:"EmailNotSent"`
	WhatsappSent		string `json:"WhatsappSent" 		db:"WhatsappSent"`
	WhatsappNotSent		string `json:"WhatsappNotSent" 		db:"WhatsappNotSent"`
	IgnoreList			string `json:"IgnoreList" 			db:"IgnoreList"`
	Confirmed			string `json:"Confirmed" 			db:"Confirmed"`
	Unconfirmed			string `json:"Unconfirmed" 			db:"Unconfirmed"`
	ConfirmedByEmail	string `json:"ConfirmedByEmail" 	db:"ConfirmedByEmail"`
	UnconfirmedByEmail	string `json:"UnconfirmedByEmail" 	db:"UnconfirmedByEmail"`
	ConfirmedBySms		string `json:"ConfirmedBySms" 		db:"ConfirmedBySms"`
	UnconfirmedBySms	string `json:"UnconfirmedBySms" 	db:"UnconfirmedBySms"`
	UniqueContacts		string `json:"UniqueContacts" 		db:"UniqueContacts"`
}


