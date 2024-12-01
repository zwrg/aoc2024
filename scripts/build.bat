@echo off
set OUTPUT_DIR=bin
set BINARY_NAME=get_day.exe

if not exist %OUTPUT_DIR% (
    mkdir %OUTPUT_DIR%
)

echo Building...
go build -o %OUTPUT_DIR%\%BINARY_NAME% get_day.go
echo Build complete. Binary available at %OUTPUT_DIR%\%BINARY_NAME%