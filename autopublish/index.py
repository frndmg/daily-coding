import os
import re

from flask import Flask
from flask import request

import telegram


BOT_TOKEN = os.getenv("BOT_TOKEN")
DAILY_CODING_TOKEN = os.getenv("DAILY_CODING_TOKEN")
DAILY_CODING_ADMIN = os.getenv("DAILY_CODING_ADMIN")
DAILY_CODING_CHANNEL = os.getenv("DAILY_CODING_CHANNEL")


app = Flask(__name__)

bot = telegram.Bot(token=BOT_TOKEN)


email_re = re.compile(
    r"This problem was asked by (?P<company>\S+)\.\n\n(?P<text>(?:.*?)(?=\n{3}-{3,}))", re.DOTALL)


@app.route('/new_email', methods=['POST'])
def new_email():
    token = request.headers.get('SECURE-TOKEN')
    if token != DAILY_CODING_TOKEN:
        return ('You have no access here', 403)

    email = request.get_data().decode()

    m = email_re.search(email)
    if m is None:
        bot.send_message(chat_id=DAILY_CODING_ADMIN,
                         text="cant parse message {}".format(email))
        return

    company = m.group('company')
    text = m.group('text')
    text = re.sub(r'(?<=.)\n(?=.)', ' ', text)

    msg = """This problem was asked by #{}

{}""".format(company, text)

    bot.send_message(chat_id=DAILY_CODING_CHANNEL, text=msg)

    return ('', 200)
