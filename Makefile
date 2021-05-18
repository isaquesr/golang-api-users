GOFILES = $(shell find . -name '*.go')

default: build

workdir:
	mkdir -p workdir

build: workdir/svc-users-go

build-native: $(GOFILES)
	go build -o workdir/svc-users-go .

workdir/svc-users-go: $(GOFILES)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o workdir/svc-users-go .