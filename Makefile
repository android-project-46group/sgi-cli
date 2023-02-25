BINARY_NAME=sgi
CURRENT_REVISION = $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS = "-s -w -X main.revision=$(CURRENT_REVISION)"

EXTERNAL_TOOLS := \
	golang.org/x/pkgsite/cmd/pkgsite@latest # latest は go 1.19 以上が必要: https://github.com/golang/pkgsite#requirements

.PHONY: build help bootstrap godoc clean
.DEFAULT_GOAL := help

build:	## バージョン値にリビジョンを埋め込んでビルド。
	go build -o ${BINARY_NAME} -ldflags=$(BUILD_LDFLAGS) cmd/sgi/*

help:	## https://postd.cc/auto-documented-makefile/
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

bootstrap: ## 外部ツールをインストールする。
	for t in $(EXTERNAL_TOOLS); do \
		echo "Installing $$t ..." ; \
		go install $$t ; \
	done

godoc:	## godoc をローカルで表示する。http://localhost:8080/{module_name}
	@echo see: http://localhost:8080/github.com/android-project-46group/sgi-cli
	pkgsite

clean:
	@$(RM) ${BINARY_NAME}
