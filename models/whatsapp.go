package models

type Whatsapp struct {
	Id int64 `json:"id"`    			// Id cast as int64
    Number string `json:"number"`  		// Number cast as string
    Message string `json:"message"`  	// Message to delivery
}