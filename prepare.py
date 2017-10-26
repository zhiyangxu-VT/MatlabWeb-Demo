import sys
import platform
import subprocess

print("Checking python version")
if sys.version_info[0] != 2 or sys.version_info[1] != 7:
    print("Python 2.7 is required")
    sys.exit()

print("Looking for MATLAB")
os_type = platform.system()
matlab_engin = '\\extern\\engines\\python'
matlab_root = ''
if os_type == 'Windows':
    import os
    paths = os.getenv("PATH").split(';')
    matlab_bin = ''
    for path in paths:
        if 'MATLAB' in path:
            matlab_bin = path
            break
    
    if matlab_bin == '':
        print("Matlab is not installed")
        sys.exit()

    matlab_root = matlab_bin.rsplit('\\', 1)[0]
else:
    matlab_root = subprocess.call(["matlab", "-e"])

print("Checking python libraries")
try:
    import pip
except:
    print("Pip for python is required")
    sys.exit()

try:
    import flask
    flask_status = 0
except:
    print("Installing flask")
    if pip.main(['install', 'flask']) != 0:
        print('Errow while setting up the enviroment. Plase run as administrator')
        sys.exit()
try:
    import werkzeug.utils
    werkzeug_status = 0
except:
    print("Installing werkzeug")
    if pip.main(['install', 'werkzeug']) != 0:
        print('Errow while setting up the enviroment. Plase run as administrator')
        sys.exit()

try:
    import matlab.engine
except:
    print("Installing MATLAB python engine")
    cwd = matlab_root + matlab_engin
    cmd = 'python setup.py install'
    install_process = subprocess.Popen(cmd, cwd=cwd)
    if install_process.wait() != 0:
        print('Errow while setting up the enviroment. Plase run as administrator')
        sys.exit()

print("Demo enviroment is ready")