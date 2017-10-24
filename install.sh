#!/bin/bash

echo "looking for python2.7..."
if ! type python > /dev/null; then
    echo "python not found";
    exit; 
fi

python_major=`python -c 'import sys; print(sys.version_info[0])'`;
python_minor=`python -c 'import sys; print(sys.version_info[1])'`;
echo "python$python_major.$python_minor is installed";

if [ $python_major != '2' ] || [ $python_minor != '7' ]; then
    echo "But python2.7 is required";
    exit;
fi

echo "looking for MATLAB...";

matlab_root='no'
if ! type matlab > /dev/null; then
    echo "MATLAB command not found under $PATH";
	read -p "Manually set MATLAB root directory if you have MATLAB installed, use no to quit: " matlab_root;
else
	matlab_root=`matlab -e | sed -n -e 's/MATLAB=//p'`;
fi

if [ "$matlab_root" == 'no' ]; then
	exit;
fi
	
echo "MATLAB found at $matlab_root";

echo "checking and installing python libraries"

test_result=`python -c 'import flask' 2>&1`
if [ "$test_result" != '' ]; then
	echo "Flask is not installed, installing..."
	#`sudo pip install flask`
fi
echo "Flask is installed"

test_result=`python -c 'import werkzeug.utils' 2>&1`
if [ "$test_result" != '' ]; then
	echo "Werkzeug is not installed, installing..."
	#`sudo pip install werkzeug`
fi
echo "Wekzeug is installed"

test_result=`python -c 'import matlab.engine' 2>&1`
if [ "$test_result" != '' ]; then
	echo "Installing MATLAB engine for python"
	`cd $matlab_root/extern/engines/python/`
	`sudo python setup.py install`
fi
echo "MATLAB engine for python is installed"