pb:
	protoc --proto_path=./proto \
    	--go_out=plugins=grpc:./proto \
    	./proto/*.proto
micro:
	 protoc --proto_path=./proto \
    	--micro_out=./proto \
    	--go_out=./proto \
    	./proto/*.proto
build:
	set GOARCH=amd64
	set GOOS=linux
	go build main.go
run:
	./consignment-service