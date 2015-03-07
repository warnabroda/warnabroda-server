package routes

import (
	"encoding/json"
	"strings"
	"flag"
	"os"
	"fmt"
	
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"bitbucket.org/hbtsmith/warnabrodagomartini/messages"
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

// For now all due verifications regarding send rules is done previewsly, here we just async the e-mail send of the warn
//Deploys the message to be sent into an email struct, call the service and in case of successful send, update the warn as sent.
func SendWhatsappWarning(entity *models.Warning, db gorp.SqlExecutor){

	subject 	:= messages.GetLocaleMessage(entity.Lang_key,"MSG_HEADER_WHATSAPP")
	message 	:= SelectMessage(db, entity.Id_message)
	footer  	:= messages.GetLocaleMessage(entity.Lang_key,"MSG_FOOTER_WHATSAPP")
	
	whatsMsg 	:= models.Whatsapp {
		Id: entity.Id,
		Number: strings.Replace(entity.Contact, "+", "", 1),
		Message: subject + " :\r\n \r\n"+message.Name + ". \r\n \r\n"+footer,
	}
	fmt.Println(whatsMsg)
	sendWhatsapp(&whatsMsg)	
}

func SendWhatsappIgnoreRequest(entity *models.Ignore_List, db gorp.SqlExecutor){

	message 	:= strings.Replace(messages.GetLocaleMessage(entity.Lang_key,"MSG_SMS_IGNORE_CONFIRMATION_REQUEST"), "{{url}}", entity.Confirmation_code, 1)
	footer  	:= messages.GetLocaleMessage(entity.Lang_key,"MSG_FOOTER")

	whatsMsg 	:= models.Whatsapp {
		Id: -1*entity.Id,
		Number: strings.Replace(entity.Contact, "+", "", 1),
		Message: message + "... "+footer,
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