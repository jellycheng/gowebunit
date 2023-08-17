.PHONY: all build run clean help go_build

# 定义变量
OUTPUTDIR="./output/"
BINARY="gowebunit"
LOG_DIR="/data/logs/app/"

all: clean build

build:
	go build -o ${OUTPUTDIR}${BINARY} main.go
	cp .env.example ${OUTPUTDIR}/.env

go_build: clean
	export GO111MODULE=on
	export GOPROXY=https://goproxy.cn
	env GOOS=linux GOARCH=amd64 go build -o ${OUTPUTDIR}${BINARY} main.go

devrun:
	@go run ./main.go -config .env.example

clean:
	@if [ -f ${OUTPUTDIR}${BINARY} ] ; then rm ${OUTPUTDIR}${BINARY} ; fi

pm2_start_web:
	pm2 start "${OUTPUTDIR}${BINARY} -config=$(CONF_NAME) " --name=${BINARY} -o ${LOG_DIR}${BINARY}/pm2/out.log -e ${LOG_DIR}${BINARY}/pm2/err.log

pm2_restart_web:
	pm2 restart ${BINARY}

pm2_stop_web:
	pm2 stop ${BINARY}

pm2_delete_web:
	pm2 delete ${BINARY}

pm2_start: pm2_start_web
pm2_restart: pm2_restart_web
pm2_stop: pm2_stop_web
pm2_delete: pm2_delete_web

#显示命令帮助，如 make help
help:
	@echo "make - 删除原二进制文件后再编译生成新的二进制文件-适用于dev环境"
	@echo "make build - 本地开发环境编译 Go 代码, 生成二进制文件-适用于dev环境"
	@echo "make go_build - 编译Go代码-适用于linux环境的代码打包"
	@echo "make devrun - 直接运行 Go 代码-适用于dev环境"
	@echo "make clean - 移除二进制文件"
	@echo "CONF_NAME=/data/www/gowebunit/output/.env make pm2_start - 启动web服务"
	@echo "make pm2_stop - 停止服务"
	@echo "make pm2_restart - 重启服务"
