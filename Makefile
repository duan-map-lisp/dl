# 文件名称: Makefile
# 文件功能: Makefile入口
# 该版作者: DXY

# 中国需要代理访问，非中国用户可以注释这里
export GOPROXY=https://goproxy.cn
SHELL:=/bin/bash

# 可执行程序
cmd := dl

auto:
	make clean
	make generate
	make all
	make test

generate:
	go build -o bin/generate cmd/generate.go
	for FileName in `find -name '*_generate_test.go'`; \
	do \
		echo "./bin/generate $$FileName"; \
		./bin/generate $$FileName; \
	done
	go fmt >& /dev/null; \
	true;

all:
	for CmdName in $(cmd); \
	do \
		echo "go build -o ./bin/$$CmdName ./cmd/$$CmdName.go"; \
		go build -o ./bin/$$CmdName ./cmd/$$CmdName.go; \
	done

test:
	for FileName in `find ./test -name '*.json'`; \
	do \
		echo "./bin/dl $$FileName"; \
		./bin/dl $$FileName; \
	done

clean:
	rm -rfdv ./bin/*

dist_clean: clean
	rm -rfdv ./lib/*
	rm -rfdv `find -name '*_generate_drop.go'`

.PHONY: \
	auto \
	all \
	test \
	clean \
	dist_clean \
	generate \
	build
