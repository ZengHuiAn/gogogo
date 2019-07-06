#!/usr/bin/env bash

arg1=$1

if [[ ! -n "$arg1" ]]; then
     echo " build success args is null"
     exit
fi

if [[ "${arg1}" == "pb" ]]; then
    protoc --proto_path=./proto \
    	--go_out=plugins=grpc:./proto \
    	./proto/*.proto
elif [[ "${arg1}" == "micro" ]]; then
    protoc --proto_path=./proto \
    --micro_out=./proto \
    --go_out=./proto \
    ./proto/*.proto
fi
 echo " build success ${arg1}"



