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
except:
    print("Installing flask")
    pip.main(['install', 'flask'])

try:
    import werkzeug.utils
except:
    print("Installing werkzeug")
    pip.main(['install', 'werkzeug'])

try:
    import matlab.engine
except:
    print("Installing MATLAB python engine")
    cwd = matlab_root + matlab_engin
    cmd = 'python setup.py install'
    install_process = subprocess.Popen(cmd, cwd=cwd)
    install_process.wait()

print("Demo enviroment is ready")