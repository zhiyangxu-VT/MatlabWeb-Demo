@echo off

cd ".\Matlab Side\"

echo this demo runs the server listening on 0.0.0.0 and handling requests from any domines
python matlab_handler.py -w all -l 0.0.0.0
pause