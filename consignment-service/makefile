pb:
	protoc --proto_path=./proto \
    	--go_out=plugins=grpc:./proto \
    	./proto/*.proto
micro:
	 protoc --proto_path=./proto \
    	--micro_out=./proto \
    	--go_out=./proto \
    	./proto/*.proto