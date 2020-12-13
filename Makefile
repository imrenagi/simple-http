build:
	CGO_ENABLED=0 GOOS=linux go build -a -o bin/video ./video/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -o bin/www ./www/main.go