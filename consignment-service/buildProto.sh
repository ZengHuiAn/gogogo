#!/usr/bin/env bash
    protoc --proto_path=./proto \
	--micro_out=./proto \
	--go_out=./proto \
	./proto/*.proto

