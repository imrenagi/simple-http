build:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -a -o bin/video ./video/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -a -o bin/www ./www/main.go
	chmod a+x bin/video
	chmod a+x bin/www