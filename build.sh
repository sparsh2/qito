set -exuo pipefail

test() {
    go fmt ./...
    go test ./...
}

build() {
    python3 ./build.py
}

test
build
