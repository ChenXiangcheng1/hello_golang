all: help
help:
	@echo "make all	默认"
	@echo "make go.mod	初始化go.mod"
	@echo "make run	运行base/main.go"
	@echo "make help	帮助"
test_config:  # 测试.golangci.yml配置
	@golangci-lint -v linters
go.mod:  # 初始化go.mod
	@rm -f go.mod
	@go mod init github.com/chenxiangcheng1/hello-golang >/dev/null 2>&1
	@go mod tidy
run: go.mod  # 运行base/main.go
	@go run ./base/main.go
