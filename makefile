pb:
	cd consignment-service && sh buildProto.sh pb
micro:
	cd consignment-service && sh buildProto.sh micro
build:
	make pb
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .

