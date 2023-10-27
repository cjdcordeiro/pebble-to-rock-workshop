from flask import Flask
from markupsafe import escape
import datetime

HOST = "0.0.0.0"
PORT = 80

app = Flask(__name__)


@app.route("/<path:path>")
def hello(path):
    if not path.startswith("/"):
        path = "/" + path
    message = """
        <html>
            <head>
                <title>Cool ROCKServer!</title>
            </head>
            <body>
                <p>Server time: {server_time}</p>
                <p>Path: {path}</p>
                <hr/>
                <p>
                    Hi, rockstar! Your ROCK is working well!
                    <a href="https://youtu.be/dQw4w9WgXcQ">Learn more.</a>
                </p>
            </body>
        </html>
    """.format(
        server_time=datetime.datetime.now(),
        path=escape(path),
    )
    return message


@app.route("/")
def index():
    return hello("/")


if __name__ == "__main__":
    app.run(host=HOST, port=PORT)
