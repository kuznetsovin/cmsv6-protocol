all: deploy

deploy:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cmsv6-srv
	scp cmsv6-srv root@138.201.190.210:/opt