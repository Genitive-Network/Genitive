########## build ##########
.Phony: build
build:
	GOOS=linux GOARCH=amd64 go build -o dist/server ./cmd/main.go
