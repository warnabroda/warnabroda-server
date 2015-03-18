package models


// DefaultStruct should be used for any given situation where all what is needed is a Simple Id/Name Struct
type DefaultStruct struct {
    Id int64 `json:"id"`    			// Id cast as int64
    Name string `json:"name"`  			// Name cast as string
    Lang_key string `json:"lang_key"`  	// Lang_key language used at Name field
    Active bool `json:"active"`
}

type MessageStruct struct {
	Id int64 `json:"id"`    			// Id cast as int64
    Name string `json:"name"`  			// Name cast as string
    Lang_key string `json:"lang_key"`  	// Lang_key language used at Name field
    Active bool `json:"active"`
    Last_modified_by int  `json:"Last_modified_by"`	
    Last_modified_date string `json:"Last_modified_date"`
}
