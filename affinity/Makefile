.PHONY: all

# compile file
BUILD_FILE=main.go

# compile build result
TARGET_NAME=affinity

# support OS
# mac os "darwin"
# linux "linux"
# windows "windows"
GOOS=linux

# compile platform, default amd64
GOARCH=amd64

all: format clean build

format:
	gofmt -w .

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-s -w" -v -o $(TARGET_NAME)

clean:
	rm -fr $(TARGET_NAME)
