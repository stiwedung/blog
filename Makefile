.PHONY: build clean

export GOPROXY=https://goproxy.cn
export GO111MODULE=on

ifeq ($(shell uname), Linux)
	OUTPUTSUFFIX=
else
	OUTPUTSUFFIX=.exe
endif

build:
	@if [ ! -d bin ]; then mkdir bin; fi
	go build -o bin/blog$(OUTPUTSUFFIX) ./cmd
	@echo "build successfully!"

clean:
	rm bin/blog$(OUTPUTSUFFIX)