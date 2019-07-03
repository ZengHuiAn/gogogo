#!/usr/bin/env bash
echo "第一个参数为：$1";

if [[ $# == 1 ]];then
    go run "$1/main.go"
else
    echo " args count is error"
fi