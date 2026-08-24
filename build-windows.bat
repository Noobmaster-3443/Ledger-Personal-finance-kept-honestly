@echo off
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-H windowsgui" -o Ledger.exe .
echo.
echo Built Ledger.exe
pause
