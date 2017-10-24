@ECHO off

REM goto python_check
goto matlab_check

:python_check
ECHO looking for python2.7...

FOR /F %%i in ('python -c "import sys; print(sys.version_info[0])"') do set python_major=%%i
FOR /F %%i in ('python -c "import sys; print(sys.version_info[1])"') do set python_minor=%%i

IF NOT "%python_major%" == "2" goto no_python
IF NOT "%python_minor%" == "7" goto no_python

ECHO python2.7 found
goto end

:matlab_check
ECHO looking for matlab_check
goto end

FOR /F %%i in ('matlab -nosplash -nodesktop -minimize -r "disp(matlabroot)"') do set matlab_root=%%i
ECHO $Env:PATH
ECHO matlab_root
goto end


:no_python
ECHO python2.7 is not found
goto :end

:end