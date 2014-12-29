package models

// This struct represents all data needed to send an e-mail, it contains all required data from Mandrill's API + sending data.
type Email struct {
      
    MandrillKey string      // Mandrill Key configured over mandrilapp.com
  
    TemplatePath string     // The path to the template that should be used!
  
    Subject string          // E-mail's subject
  
    Content string          // When @TemplatePath is not used Content should be considered
  
    ToAddress string        // E-mail's destination address
  
    FromAddress string      // Mostly warnabroda@gmail.com
  
    FromName string         // Mostly "Warn A Broda" 
      
    LangKey string          // Language the subject and content should use 

    Async bool              // True if mandrill's API should send it asynchronously

    UseContent bool         // True if Mandrill's API should consider @Content over @TemplatePath 

    HTMLContent bool        // True if Mandrill's API should consider @Content as HTML
}