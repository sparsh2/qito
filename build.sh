set -exuo pipefail

test() {
    go fmt ./...
    go test ./...
}

build() {
    go build -o ./target/psm
}

test
build
