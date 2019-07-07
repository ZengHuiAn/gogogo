#!/usr/bin/env bash

echo "第一个参数为：$1";

if [[ $# == 1 ]];then
    docker-compose  -f  "$1.yml" up
else
    echo " args count is error"
fi


function pause(){
        read -n 1 -p "$*" INP
        if [[ ${INP} != '' ]] ; then
                echo -ne '\b \n'
        fi
}

pause 'Press any key to continue...'