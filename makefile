.DEFAULT_GOAL := build

APPLICATION=go-kikidelivery

build:
	go build -o "${APPLICATION}"

clean:
	go clean
	rm --force "cp.out"
	rm --force nohup.out

test:
	go test -v

cover:
	go test -coverprofile cp.out
	go tool cover -html=cp.out

run1:
	./"${APPLICATION}" 1
	100 3
	PKG1 5 5 OFR001
	PKG2 15 5 OFR002
	PKG3 10 100 OFR003


lint:
	golangci-lint run --enable-all
