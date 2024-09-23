.PHONY: build
build:
	GOOS=windows GOARCH=amd64 go build -o win64/touch.exe
	GOOS=windows GOARCH=386 go build -o win32/touch.exe
	GOOS=windows GOARCH=arm64 go build -o winARM/touch.exe