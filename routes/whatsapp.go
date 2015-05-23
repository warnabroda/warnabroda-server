package routes

import (
	"encoding/json"
	"strings"
	"flag"
	"os"
	//"fmt"
	
	"warnabrodagomartini/models"
	"warnabrodagomartini/messages"
	"github.com/coopernurse/gorp"	
	"github.com/jjeffery/stomp"
)

var (	
	RabbitMQUser 		= os.Getenv("WARNARABBITMQUSER")
	RabbitMQPass		= os.Getenv("WARNARABBITMQPASS")
	HostWarnabroda		= os.Getenv("WARNAHOST")
	WarnaQueueWhats 	= os.Getenv("WARNAQUEUEWHATSAPP")
	serverAddr 			= flag.String("server", HostWarnabroda+":61613", "STOMP server endpoint")
	queueName 			= flag.String("queue", WarnaQueueWhats, "Destination queue")
	helpFlag 			= flag.Bool("help", false, "Print help text")
	stop 				= make(chan bool)
)

var options = stomp.Options{
	Login:    RabbitMQUser,
	Passcode: RabbitMQPass,
	Host:     "/",
}


func init(){	
	flag.Parse()
}


func ProcessWhatsapp(entity *models.Warning, db gorp.SqlExecutor){

	var footer string
	if entity.WarnResp != nil {
		if entity.WarnResp.Id_contact_type == 1 {
			go SendEmailReplyRequestAcknowledge(entity.WarnResp, db)
		} else {
			go SendWhatsappReplyRequestAcknowledge(entity.WarnResp, db)
		}
		replyToLocaleMsg := messages.GetLocaleMessage(entity.Lang_key, "MSG_FOOTER_REPLY")
		footer = strings.Replace(replyToLocaleMsg, "{{url_reply}}", models.URL_REPLY + "/" + entity.WarnResp.Resp_hash, 1)
	}

	subject 	:= messages.GetLocaleMessage(entity.Lang_key,"MSG_HEADER_WHATSAPP")
	message 	:= SelectMessage(db, entity.Id_message)
	footer  	+= "\r\n"+messages.GetLocaleMessage(entity.Lang_key,"MSG_FOOTER_WHATSAPP")
	
	whatsMsg 	:= models.Whatsapp {
		Id: entity.Id,
		Number: strings.Replace(entity.Contact, "+", "", 1),
		Message: subject + " :\r\n \r\n"+message.Name + " \r\n \r\n"+footer,
		Type: models.MSG_TYPE_WARNING,
	}

	sendWhatsapp(&whatsMsg)	
}

func SendWhatsappReplyDone(entity *models.Warning, db gorp.SqlExecutor){

	
	message := SelectMessage(db, entity.Id_message)

	whatsappMsg := messages.GetLocaleMessage(entity.Lang_key, "MSG_SMS_HEADER") + " \r\n \r\n"

	mainBodyMsg := messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_BODY_MAIN")
	mainBodyMsg = strings.Replace(mainBodyMsg, "{{sent_msg}}", "'" + message.Name+ "'", 1) 
	mainBodyMsg = strings.Replace(mainBodyMsg, "{{contact_msg}}", "'" + entity.Contact+ "'", 1)	


	whatsappMsg += mainBodyMsg + " \r\n \r\n"
	whatsappMsg += messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_BODY_LINK") + " \r\n \r\n"
	whatsappMsg += models.URL_REPLY + "/" + entity.WarnResp.Read_hash
	
	whatsMsg 	:= models.Whatsapp {
		Id: entity.WarnResp.Id,
		Number: strings.Replace(entity.WarnResp.Reply_to, "+", "", 1),
		Message: whatsappMsg, 
		Type: models.MSG_TYPE_REPLY,
	}
	
	sendWhatsapp(&whatsMsg)	
}

func SendWhatsappReplyRequestAcknowledge(entity *models.WarningResp, db gorp.SqlExecutor){
	
	
	senderLocaleMsg := messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_SENDER_MSG")
	senderLocaleMsg = strings.Replace(senderLocaleMsg, "{{reply_to}}", entity.Reply_to, 1)
	replyToWhatsMsg := models.Whatsapp{
		Id: entity.Id,
		Number: strings.Replace(entity.Reply_to, "+", "", 1),
		Message: senderLocaleMsg,
		Type: "warning-reply",
	}
	
	sendWhatsapp(&replyToWhatsMsg)	

}

func SendWhatsappIgnoreRequest(entity *models.Ignore_List, db gorp.SqlExecutor){

	message 	:= strings.Replace(messages.GetLocaleMessage(entity.Lang_key,"MSG_SMS_IGNORE_CONFIRMATION_REQUEST"), "{{url}}", entity.Confirmation_code, 1)
	footer  	:= messages.GetLocaleMessage(entity.Lang_key,"MSG_FOOTER")

	whatsMsg 	:= models.Whatsapp {
		Id: -1*entity.Id,
		Number: strings.Replace(entity.Contact, "+", "", 1),
		Message: message + "... "+footer,
		Type: models.MSG_TYPE_IGNORE,
	}

	sendWhatsapp(&whatsMsg)
}

func sendWhatsapp(entity *models.Whatsapp){
	entityJson,_ := json.Marshal(entity)	
	body 		:= string(entityJson[:])	
	
	go sendMessages(body)
	
	<-stop
}

func sendMessages(msg string) {
	defer func() {
		stop <- true
	}()	

	conn, err := stomp.Dial("tcp", *serverAddr, options)
	if err != nil {
		println("cannot connect to server", err.Error())
		return
	}
	err = conn.Send(*queueName, "text/plain",[]byte(msg), nil)
	if err != nil {
		println("failed to send to server", err)
		return
	}


}