@echo off

set BINARY=bin\get_day.exe
set BUILD_SCRIPT=scripts\build.bat

if not exist "%BINARY%" (
    echo Binary not found. Running build script...
    if not exist "%BUILD_SCRIPT%" (
        echo Build script not found. Exiting...
        exit /b 1
    )
    call "%BUILD_SCRIPT%"
)

if exist "%BINARY%" (
    echo Running %BINARY% with parameter: %1
    "%BINARY%" %1
) else (
    echo Build failed or binary not found. Exiting...
    exit /b 1
)
