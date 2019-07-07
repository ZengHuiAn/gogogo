#!/usr/bin/env bash

arg1=$1
function pause(){
        read -n 1 -p "$*" INP
        if [[ ${INP} != '' ]] ; then
                echo -ne '\b \n'
        fi
}
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
    --go_out=plugins=micro:./proto \
    ./proto/*.proto
fi
 echo " build success ${arg1}"

 #使用时：
pause 'Press any key to continue...'



