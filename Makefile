release:
	GOOS=darwin GOARCH=amd64 go build -o todo-mac
	GOOS=linux GOARCH=amd64 go build -o todo-linux
	GOOS=windows GOARCH=amd64 go build -o todo-windows
