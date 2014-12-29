package models

// Represents all data required to send a SMS via MobilePronto, this struct is totally binded to the SMS sender
type SMS struct {
      
    CredencialKey string	// Credential key registered at MobilePronto
  
    Content string			// SMS content

    URLPath string			// MobilePronto URL used to SEND the SMS and to parse MobilePronto Response 
  
    Scheme string    		// http or https
  
    Host string				// MobilePronto Domain
  
    Project string			// Warnabroda
  
    AuxUser string			// dummy value, I dont know what is this for.
      
    MobileNumber string		// Destination Cell Number

    SendProject string		// Whether it should send the project name along with SMS or not.

}