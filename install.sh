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

echo "looking for matlab..."
if ! type matlab > /dev/null; then
    echo "MATLAB not found";
    exit;
fi

matlab_root=`matlab -e | sed -n -e 's/MATLAB=//p'`
echo "MATLAB found at $matlab_root"

echo "checking and installing python libraries"
    echo `python -c 'import flas'`
    echo `python -c 'import werkzeug.util'`
    echo `python -c 'import matlab.engin'`