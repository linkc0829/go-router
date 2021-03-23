@echo off
if not exist bin (
    mkdir bin
)
for /f "usebackq" %%i in (`dir /b /on /a:d .\cmd`) do (
    echo %%i
    go build -o ./bin/%%i.exe ./cmd/%%i
)