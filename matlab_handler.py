#! python

import time
import os
import json

from flask import Flask
from flask import Response
from flask import request
from werkzeug.utils import secure_filename

import matlab.engine

app = Flask(__name__)

@app.route("/")
def hello():
    return "Hello World!"

@app.route('/matlab/<action>', methods=['POST'])
def handle_matlab(action):
    saved_files = save_upload_files(request.files)
    result = matlab_analyse(action, saved_files)

    resp = Response(result)
    resp.headers['Access-Control-Allow-Origin'] = '*'
    
    return resp 
    
def save_upload_files(uploadfile):
    a_file = uploadfile['uploadfile']
    file_path = 'uploads'
    if not os.path.exists(file_path):
        os.makedirs(file_path)

    time_stamp = str(time.time())
    file_name = time_stamp + '-' + secure_filename(a_file.filename)
    
    temp_file = os.path.join(file_path, file_name)
    a_file.save(temp_file)

    return temp_file

def matlab_analyse(action, files):
    eng = matlab.engine.start_matlab()
    matlab_func = getattr(eng, action)
    result = matlab_func(files)

    print(result)
    
    return json.dumps(result)

if __name__ == "__main__":
    app.run()