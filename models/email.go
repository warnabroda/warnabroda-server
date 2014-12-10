package models

type Email struct {
      
    MandrillKey string
  
    TemplatePath string
  
    Subject string  
  
    Content string
  
    ToAddress string
  
    FromAddress string
  
    FromName string
      
    LangKey string

    Async bool

    UseContent bool

    HTMLContent bool
}