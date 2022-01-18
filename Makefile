Version := $(shell git describe --tags --dirty)
GitCommit := $(shell git rev-parse HEAD)
#LDFLAGS := "-s -w -X github.com/willfore/mattermost_to_slack/cmd.Version=$(Version) -X github.com/willfore/mattermost_to_slack/cmd.GitCommit=$(GitCommit)"
PLATFORM := $(shell ./hack/platforms.sh)
SOURCE_DIRS = cmd main.go
export GO111MODULE=on

.PHONY: all
all: gofmt test build dist hash

.PHONY: build
build:
	go build

.PHONY: gofmt
gofmt:
	@test -z $(shell gofmt -l -s $(SOURCE_DIRS) ./ | tee /dev/stderr) || (echo "[WARN] Fix formatting issues with 'make fmt'" && exit 1)

.PHONY: test
test:
	CGO_ENABLED=0 go test $(shell go list ./... | grep -v /vendor/|xargs echo) -cover

.PHONY: dist
dist:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/mm2slack
	CGO_ENABLED=0 GOOS=darwin go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/mm2slack-darwin
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/mm2slack-armhf
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/mm2slack-arm64
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/mm2slack.exe

.PHONY: hash
hash:
	rm -rf bin/*.sha256 && ./hack/get_hash.sh
