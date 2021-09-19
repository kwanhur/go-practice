.PHONY: all

# compile file
BUILD_FILE=main.go

# compile build result
TARGET_NAME=confd

# support OS
# mac os "darwin"
# linux "linux"
# windows "windows"
GOOS=darwin

# compile platform, default amd64
GOARCH=amd64

# number processors
GOMAXPROCS=2

SCHED_INTERVAL=1000


all: format clean build

format:
	gofmt -w .

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-s -w" -v -o $(TARGET_NAME)

clean:
	rm -fr $(TARGET_NAME)

debug:
    GOMAXPROCS=$(GOMAXPROCS) GODEBUG=schedtrace=$(SCHED_INTERVAL),scheddetail=1 ./$(TARGET_NAME)