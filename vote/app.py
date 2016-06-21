from flask import Flask, render_template, request, make_response
import socket
from wsgiref.handlers import CGIHandler
import os
import random
import dockerrun

option_a = os.getenv('OPTION_A', "Cats")
option_b = os.getenv('OPTION_B', "Dogs")
hostname = socket.gethostname()

app = Flask(__name__)

client = dockerrun.from_env()

@app.route("/", methods=['POST','GET'])
def hello():
    voter_id = request.cookies.get('voter_id')
    if not voter_id:
        voter_id = hex(random.getrandbits(64))[2:-1]

    vote = None

    if request.method == 'POST':
        vote = request.form['vote']
        client.run(
            "bfirsh/serverless-record-vote-task",
            [voter_id, vote],
            detach=True,
            network_mode="serverlessdockervotingapp_default"
        )

    resp = make_response(render_template(
        'index.html',
        option_a=option_a,
        option_b=option_b,
        hostname=hostname,
        vote=vote,
    ))
    resp.set_cookie('voter_id', voter_id)
    return resp


if __name__ == "__main__":
    CGIHandler().run(app)
