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
from stomp.listener import ConnectionListener


class WhatsappStompClientListener(ConnectionListener):

    def __init__(self):
        self.credentials = None
        self.whatsapp = WhatsappCli(self.credentials)


    def on_message(self, headers, body):
        msg = json.loads(body)
        if "destination" in body and "message" in body:
            dest = body["destination"]
            msg = body["message"]
            sent = self.get_whatsap().send_msg(dest, msg)
            LOG.debug("MSG SENT %s", sent)


    def get_whatsap(self):
        if not self.whatsapp:
            self.whatsapp = WhatsappCli(self.get_credentials())
        return self.whatsapp

    def get_credentials(self):
        if not self.credentials:
            self.credentials = (os.environ["WARNA_WHATSAPP_NUMBER"], os.environ["WARNA_WHATSAPP_PASS"])
        return self.credentials



if __name__ == "__main__":
    logging.basicConfig()
    conn = stomp.Connection()
    conn.set_listener('', WhatsappStompClientListener())
    conn.start()
    conn.connect()

    conn.subscribe(destination='/queue/whatsapp', id=1, ack='auto')




