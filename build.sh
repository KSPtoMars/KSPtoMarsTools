#!/bin/bash
#

GOPATH=/home/darko/Documents GOOS=windows GOARCH=386 CGO_ENABLED=1 CXX_FOR_TARGET=i686-w64-mingw32-g++ CC_FOR_TARGET=i686-w64-mingw32-gcc /usr/local/go/bin/go build modpack.go
