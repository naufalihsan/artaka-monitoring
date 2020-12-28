set -xe

GOARCH=amd64 GOOS=linux go build -o bin/application application.go