#!/bin/sh

if [ ! -f start.sh ]; then
    echo "start.sh must run from its folder"
    exit 1
fi

if [ $(uname) = Linux ];then
    exe=./bin/blog
else
    exe=./bin/blog.exe
fi

if [ ! -f ${exe} ];then
    echo "start build project"
    make
    if [ ! -f ${exe} ];then
        echo "build project failed"
        exit 1
    fi
fi

${exe}
