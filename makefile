# SwaggerExp version
VERSION=$(shell cat pkg/runner/banner.go |grep "const Version ="|cut -d"\"" -f2)
# Output File Location
DIR=data/v${VERSION}
$(shell mkdir -p ${DIR})

# Go build flags
LDFLAGS=-ldflags "-s -w"

default:
	go build ${LDFLAGS} -o ${DIR}/FuckFingerprint cmd/main.go

# Compile Server - Windows x64
windows:
	export GOOS=windows;export GOARCH=amd64;go build ${LDFLAGS} -o ${DIR}/FuckFingerprint-Windows-x64.exe cmd/main.go

# Compile Server - Linux x64
linux:
	export GOOS=linux;export GOARCH=amd64;go build ${LDFLAGS} -o ${DIR}/FuckFingerprint-Linux-x64 cmd/main.go

# Compile Server - Darwin x64
darwin:
	export GOOS=darwin;export GOARCH=amd64;go build ${LDFLAGS} -o ${DIR}/FuckFingerprint-Darwin-x64 cmd/main.go

# clean
clean:
	rm -rf ${DIR}