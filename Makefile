.DEFAULT_GOAL := all

MKDIR_P = mkdir -p
OUT_DIR = ./out
BUILD_DIR = ./build

${OUT_DIR}:
	${MKDIR_P} ${OUT_DIR}

${BUILD_DIR}:
	${MKDIR_P} ${BUILD_DIR}

${BUILD_DIR}/bin/golangci-lint: ${BUILD_DIR}
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "${BUILD_DIR}/bin" v1.27.0

deps:
	go mod download

lint: ${BUILD_DIR}/bin/golangci-lint
	${BUILD_DIR}/bin/golangci-lint run

test: ${OUT_DIR}
	go test -v -coverprofile=out/coverage.profile ./...

coverage: lint ${OUT_DIR}/coverage.profile
	go tool cover -html out/coverage.profile -o out/coverage.html

all: deps lint test coverage

clean:
	rm -rf out/coverage.html
	rm -rf out/coverage.profile

${OUT_DIR}/coverage.profile: test
${OUT_DIR}/coverage.html: coverage
