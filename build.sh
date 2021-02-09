GOOS=linux GOARCH=mipsle GOMIPS=softfloat CGO_ENABLED=0 go build main.go
sz main
rm -rf main
