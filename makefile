default: build
BIN_DIR=/usr/local/bin/
BUILD_DATE=$(shell date +%Y.%m.%d)
BUILD_VER=0.0.1
GIT_COMMIT=$(shell git rev-parse --short HEAD)
BUILD_CFG=debug

build:
	env GOOS=darwin go build \
	-ldflags "-X 'main.buildDate=$(BUILD_DATE)' -X 'main.buildVersion=$(BUILD_VER)' -X 'main.buildCommit=$(GIT_COMMIT)' -X 'main.buildConfig=$(BUILD_CFG)'" \
	-o $(BIN_DIR) ./...

test:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

install:
	go install .

run:
	./$(BIN_DIR)/godog