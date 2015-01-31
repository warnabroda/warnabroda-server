import os
import logging

from yowsup.demos import sendclient

LOG = logging.getLogger(__name__)
LOG.setLevel(logging.DEBUG)

class WhatsappCli():

    def __init__(self, credentials):
        self.credentials = credentials

    def send_msg(self, dest, msg, encrypted=False):
        sent = False
        try:
            stack = sendclient.YowsupSendStack(self.credentials, [(dest, msg)], encrypted)
            stack.start()
        except KeyboardInterrupt:
            sent = True
        except Exception:
            LOG.exception("Could not send msg %s to %s - with encryption (%s)", msg, dest, encrypted)

        return sent



import simplejson as json
import stomp
import requests

class WhatsappStompClientListener(object):

    def __init__(self):
        self.credentials = None
        self.whatsapp = None


    def on_message(self, headers, message):
        try:
            message = json.loads(message)
            if "number" in message and "message" in message:
                id = message["id"]
                dest = message["number"]
                msg = message["message"]
                sent = self.get_whatsap().send_msg(dest, msg)
                LOG.debug("MSG SENT %s", sent)
                self.send_confirmation(id, sent)
        except:
            LOG.info("Could not send message %s", message)


    def get_whatsap(self):
        if not self.whatsapp:
            self.whatsapp = WhatsappCli(self.get_credentials())
        return self.whatsapp

    def get_credentials(self):
        if not self.credentials:
            self.credentials = (os.environ["WARNA_WHATSAPP_NUMBER"], os.environ["WARNA_WHATSAPP_PASS"])
        return self.credentials

    def send_confirmation(self, msg_id, msg):
        data = {"id":msg_id, "name":msg}
        requests.post(url="www.warnabroda.com:3000/warnabroda/warning-confirm", data=data)
        LOG.debug("Post sent")


if __name__ == "__main__":
    print "Iniciando cliente"
    logging.basicConfig()
    conn = stomp.Connection()
    conn.set_listener('', WhatsappStompClientListener())
    conn.start()
    print "Connectando"
    conn.connect(os.environ["WARNARABBITMQUSER"], os.environ["WARNARABBITMQPASS"])
    print "Connectado"

    print "Se inscrevendo na lista"
    conn.subscribe(destination=os.environ["WARNAQUEUEWHATSAPP"], id=1, ack='auto')
    print "Inscrito"

    try:
        while True:
            pass
    except:
        print "Vai desconectar"
    finally:
        conn.disconnect()
        print "Desconectado"





