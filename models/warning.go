package models


// Struct that represents a warning request, used to log all possible data from a sending user.
type Warning struct {
    Id int64 `json:"id"`
    
    Id_message int64 `json:"id_message"`
  
    Id_contact_type int64 `json:"id_contact_type"`
  
    Contact string `json:"contact"`
  
    Sent bool `json:"sent"`
  
    Message string `json:"message"`
  
    Ip string `json:"ip"`
  
    Browser string `json:"browser"`
  
    Operating_system string `json:"operating_system"`
  
    Device string `json:"device"`
  
    Raw string `json:"raw"`
  
    Created_by string `json:"created_by"`
  
    Created_date string `json:"created_date"`
  
    Last_modified_by string `json:"last_modified_by"`
  
    Last_modified_date string `json:"last_modified_date"`    

    Lang_key string `json:"lang_key"`

    Timezone string `json:"timezone"`

    WarnResp *WarningResp `json:"warning_resp" db:"-"`  

}


type Warn struct {
    Id int64 `json:"Id" db:"id"`
    
    Message string `json:"Message" db:"message"`
  
    ContactType string `json:"ContactType" db:"contact_type"`
  
    Contact string `json:"Contact" db:"contact"`
  
    Sent bool `json:"Sent" db:"sent"`    
  
    CreatedDate JDate `json:"CreatedDate" db:"created_date"`
  
}

type WarningResp struct {
    Id int64 `json:"id"`
    
    Id_warning int64 `json:"id_warning"`

    Id_contact_type int64 `json:"id_contact_type"`
  
    Resp_hash string `json:"resp_hash"`

    Read_hash string `json:"read_hash"`
  
    Message string `json:"message"`

    Reply_to string `json:"reply_to"`
  
    Ip string `json:"ip"`
  
    Browser string `json:"browser"`
  
    Operating_system string `json:"operating_system"`
  
    Device string `json:"device"`
  
    Raw string `json:"raw"`
  
    Created_date string `json:"created_date"`
  
    Reply_date string `json:"reply_date"`

    Response_read string `json:"response_read"`
  
    Lang_key string `json:"lang_key"`

    Timezone string `json:"timezone"`

    Sent bool `json:"sent"`
  
}