package routes

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/coopernurse/gorp"
	"github.com/jjeffery/stomp"
	"gitlab.com/warnabroda/warnabrodagomartini/messages"
	"gitlab.com/warnabroda/warnabrodagomartini/models"
)

var (
	RabbitMQUser    = os.Getenv("WARNARABBITMQUSER")
	RabbitMQPass    = os.Getenv("WARNARABBITMQPASS")
	HostWarnabroda  = os.Getenv("WARNAHOST")
	WarnaQueueWhats = os.Getenv("WARNAQUEUEWHATSAPP")
	serverAddr      = flag.String("server", HostWarnabroda+":61613", "STOMP server endpoint")
	queueName       = flag.String("queue", WarnaQueueWhats, "Destination queue")
	helpFlag        = flag.Bool("help", false, "Print help text")
	stop            = make(chan bool)
)

var options = stomp.Options{
	Login:    RabbitMQUser,
	Passcode: RabbitMQPass,
	Host:     "/",
}

func init() {
	flag.Parse()
}

func ProcessWhatsapp(entity *models.Warning, db gorp.SqlExecutor) {

	var message string
	structMsg := SelectMessage(db, entity.Id_message)
	if IsLoadedInRedis(entity.Contact) {

		message = messages.GetLocaleMessage(entity.Lang_key, "MSG_WHATSAPP_HEADER_DEFAULT") + "\r\n \r\n"
		message += "'" + structMsg.Name + ".'\r\n \r\n"

		if entity.WarnResp != nil {
			message += messages.GetLocaleMessage(entity.Lang_key, "MSG_WHATSAPP_REPLY_DEFAULT") + " " + ShortUrl(models.URL_REPLY+"/"+entity.WarnResp.Resp_hash) + "\r\n \r\n"

			processReply(entity, db)

		}

		message += strings.Replace(messages.GetLocaleMessage(entity.Lang_key, "MSG_WHATSAPP_FOOTER_DEFAULT"), "{{url_ignoreme}}", models.URL_IGNORE_REQUEST+"/"+entity.Contact, 1)

	} else {

		message = messages.GetLocaleMessage(entity.Lang_key, "MSG_WHATSAPP_HEADER_FIRSTTIMER") + "\r\n \r\n"
		message += "'" + structMsg.Name + ".'\r\n \r\n"

		if entity.WarnResp != nil {
			message += messages.GetLocaleMessage(entity.Lang_key, "MSG_WHATSAPP_REPLY_FIRSTTIMER") + " " + ShortUrl(models.URL_REPLY+"/"+entity.WarnResp.Resp_hash) + "\r\n \r\n"

			processReply(entity, db)

		}

		message += strings.Replace(messages.GetLocaleMessage(entity.Lang_key, "MSG_WHATSAPP_FOOTER_FIRSTTIMER"), "{{url_ignoreme}}", models.URL_IGNORE_REQUEST+"/"+entity.Contact, 1)

	}

	fmt.Println(message)

	AddPhoneToRedis(entity.Contact)

	whatsMsg := models.Whatsapp{
		Id:      entity.Id,
		Number:  strings.Replace(entity.Contact, "+", "", 1),
		Message: message,
		Type:    models.MSG_TYPE_WARNING,
	}

	sendWhatsapp(&whatsMsg)
}

func processReply(entity *models.Warning, db gorp.SqlExecutor) {
	if entity.Id_contact_type == 1 {
		SendEmailReplyRequestAcknowledge(entity.WarnResp, db)
	} else {
		SendWhatsappReplyRequestAcknowledge(entity, db)
	}
}

func SendWhatsappReplyDone(entity *models.Warning, db gorp.SqlExecutor) {

	message := SelectMessage(db, entity.Id_message)

	whatsappMsg := messages.GetLocaleMessage(entity.Lang_key, "MSG_SMS_HEADER") + " \r\n \r\n"

	mainBodyMsg := messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_BODY_MAIN")
	mainBodyMsg = strings.Replace(mainBodyMsg, "{{sent_msg}}", "'"+message.Name+"'", 1)
	mainBodyMsg = strings.Replace(mainBodyMsg, "{{contact_msg}}", "'"+entity.Contact+"'", 1)

	whatsappMsg += mainBodyMsg + " \r\n \r\n"
	whatsappMsg += messages.GetLocaleMessage(entity.Lang_key, "MSG_REPLY_BODY_LINK") + " \r\n \r\n"
	whatsappMsg += ShortUrl(models.URL_REPLY + "/" + entity.WarnResp.Read_hash)

	whatsMsg := models.Whatsapp{
		Id:      entity.WarnResp.Id,
		Number:  strings.Replace(entity.WarnResp.Reply_to, "+", "", 1),
		Message: whatsappMsg,
		Type:    models.MSG_TYPE_REPLY,
	}

	sendWhatsapp(&whatsMsg)
}

//"Hola, cuando el mensaje \r\n \r\n '{{msg}}' \r\n\r\n que ha enviado a \r\n\r\n {{contact}} \r\n\r\n obtener una respuesta, se le asomó aquí \r\n\r\nCódigo del Mensaje:"
func SendWhatsappReplyRequestAcknowledge(entity *models.Warning, db gorp.SqlExecutor) {

	msg := SelectMessage(db, entity.Id_message)

	replyRequest := messages.GetLocaleMessage(entity.Lang_key, "MSG_WHATSAPP_REPLY_REQUEST")
	replyRequest = strings.Replace(replyRequest, "{{msg}}", msg.Name, 1)
	replyRequest = strings.Replace(replyRequest, "{{contact}}", entity.Contact, 1) + " " + entity.WarnResp.Resp_hash[:6] + "\r\n" + messages.GetLocaleMessage(entity.Lang_key, "MSG_FOOTER")

	replyToWhatsMsg := models.Whatsapp{
		Id:      entity.Id,
		Number:  strings.Replace(entity.WarnResp.Reply_to, "+", "", 1),
		Message: replyRequest,
		Type:    "warning-reply",
	}

	sendWhatsapp(&replyToWhatsMsg)

}

func SendWhatsappIgnoreRequest(entity *models.Ignore_List, db gorp.SqlExecutor) {

	message := strings.Replace(messages.GetLocaleMessage(entity.Lang_key, "MSG_SMS_IGNORE_CONFIRMATION_REQUEST"), "{{url}}", entity.Confirmation_code, 1)
	footer := messages.GetLocaleMessage(entity.Lang_key, "MSG_FOOTER")

	whatsMsg := models.Whatsapp{
		Id:      -1 * entity.Id,
		Number:  strings.Replace(entity.Contact, "+", "", 1),
		Message: message + "... " + footer,
		Type:    models.MSG_TYPE_IGNORE,
	}

	sendWhatsapp(&whatsMsg)
}

func sendWhatsapp(entity *models.Whatsapp) {
	entityJson, _ := json.Marshal(entity)
	body := string(entityJson[:])

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
	err = conn.Send(*queueName, "text/plain", []byte(msg), nil)
	if err != nil {
		println("failed to send to server", err)
		return
	}

}
