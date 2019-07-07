pb:
	cd consignment-service && buildProto.sh pb
micro:
	cd consignment-service && buildProto.sh micro
build:
	set GOOS=linux
	cd consignment-service && make build
	cd ..
	make docker_build
docker_build:
	docker build -t consignment-service .
run:
	.\dockerRun.sh "consignment-service"
	# 在 Docker alpine 容器的 50001端口上运行 consignment-service 服务
	# 可添加 -d 参数将微服务放到后台运行
build_cli:
	set GOOS=linux
	cd consignment-cli && make build
	cd ..
	make docker_build
run_cli:
	.\dockerRun.sh "consignment-cli"

