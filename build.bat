
rem Remove already created files
del %GOPATH%\bin\linux-amd64\web-cache
del %GOPATH%\bin\windows-amd64\web-cache.exe

rem Build executable file for Linux OS
set GOOS=linux
set GOARCH=amd64
go build -o %GOPATH%\bin\linux-amd64\web-cache ./main/web_server.go

rem Build executable file for Windows OS
set GOOS=windows
set GOARCH=amd64
go build -o %GOPATH%\bin\windows-amd64\web-cache.exe ./main/web_server.go
