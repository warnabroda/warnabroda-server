package models


// Represents the database model for the ignore_list
type Ignore_List struct {
    Id int64 `json:"id"`
    
    Contact string `json:"contact"`

    Ip string `json:"ip"`
  
    Browser string `json:"browser"`
  
    Operating_system string `json:"operating_system"`
  
    Device string `json:"device"`
  
    Raw string `json:"raw"`
  
    Created_by string `json:"created_by"`
  
    Created_date string `json:"created_date"`

    Last_modified_date string `json:"last_modified_date"`

    Confirmed bool `json:"confirmed"`

    Confirmation_code string `json:"confirmation_code"`

    Message string `json:"message"`
    
}
