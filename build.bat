rem Remove ./out directory
rmdir /s/Q out

rem Build executable file for Linux OS
set GOOS=linux
set GOARCH=amd64
go build -o ./out/linux-amd64/web-cache ./main/web_server.go

rem Build executable file for Windows OS
set GOOS=windows
set GOARCH=amd64
go build -o ./out/windows-amd64/web-cache.exe ./main/web_server.go
