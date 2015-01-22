package routes

import (
	"encoding/json"
	"strings"
	"flag"
	"os"	
	
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"bitbucket.org/hbtsmith/warnabrodagomartini/messages"
	"github.com/coopernurse/gorp"	
	"github.com/jjeffery/stomp"
)

var (	
	RabbitMQUser 		= os.Getenv("WARNARABBITMQUSER")
	RabbitMQPass		= os.Getenv("WARNARABBITMQPASS")
	HostWarnabroda		= os.Getenv("WARNAHOST")
	serverAddr 			= flag.String("server", HostWarnabroda+":61613", "STOMP server endpoint")
	queueName 			= flag.String("queue", "/queue/warnabroda_whatsapp", "Destination queue")
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
func ProcessWhatsapp(warning *models.Warning, db gorp.SqlExecutor){	
	
	go sendWhatsappWarn(warning, db)
	
}

//Deploys the message to be sent into an email struct, call the service and in case of successful send, update the warn as sent.
func sendWhatsappWarn(entity *models.Warning, db gorp.SqlExecutor) {	

	subject 	:= GetRandomSubject(entity.Lang_key)
	message 	:= SelectMessage(db, entity.Id_message)
	footer  	:= messages.GetLocaleMessage(entity.Lang_key,"MSG_FOOTER")
	
	whatsMsg 	:= models.Whatsapp {
		Id: entity.Id,
		Number: strings.Replace(entity.Contact, "+", "", 1),
		Message: subject.Name + " : "+message.Name + ". "+footer,
	}
	whatsJson,_ := json.Marshal(whatsMsg)
	//fmt.Println(string(whatsJson[:]))
	body 		:= string(whatsJson[:])	

	// if err := publish(*uri, *exchangeName, *exchangeType, *routingKey, body, *reliable); err != nil {
	// 	log.Fatalf("%s", err)
	// }
	// log.Printf("published %dB OK", len(body))
	go sendMessages(body)

	//<-stop
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
	err = conn.Send(*queueName, "application/json",[]byte(msg), nil)
	if err != nil {
		println("failed to send to server", err)
		return
	}


}