set GOOS=linux
set GOARCH=amd64
go test ../... -cover
go build -o ./artifacts/app ../app.go
build-lambda-zip -o ./artifacts/app.zip ./artifacts/app