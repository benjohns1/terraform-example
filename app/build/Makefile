.PHONY: build

build:
	GOOS=linux GOARCH=amd64 go build -o ./artifacts/app ../app.go
	(cd artifacts; zip -FS app.zip ./app)
