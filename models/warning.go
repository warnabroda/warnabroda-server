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
  
}
