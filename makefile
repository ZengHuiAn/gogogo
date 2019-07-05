pb:
	cd consignment-service && sh buildProto.sh pb
micro:
	cd consignment-service && sh buildProto.sh micro
build:
	make pb
	GOOS=linux GOARCH=amd64 go build -o consignment-service consignment-service/main.go
	docker build -t consignment-service .

