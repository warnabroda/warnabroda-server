import os
import logging

from yowsup.demos import sendclient

LOG = logging.getLogger(__name__)

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

class WhatsappStompClientListener(object):

    def __init__(self):
        self.credentials = None
        self.whatsapp = WhatsappCli(self.credentials)


    def on_message(self, headers, message):
        try:
            msg = json.loads(message)
            if "destination" in message and "message" in message:
                id = message["id"]
                dest = message["number"]
                msg = message["message"]
                sent = self.get_whatsap().send_msg(dest, msg)
                LOG.debug("MSG SENT %s", sent)
                self.send_confirmation(id, sent)
        except:
            LOG.info("Could not load message %s", message)


    def get_whatsap(self):
        if not self.whatsapp:
            self.whatsapp = WhatsappCli(self.get_credentials())
        return self.whatsapp

    def get_credentials(self):
        if not self.credentials:
            self.credentials = (os.environ["WARNA_WHATSAPP_NUMBER"], os.environ["WARNA_WHATSAPP_PASS"])
        return self.credentials

    def send_confirmation(self, msg_id, msg):
        pass



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