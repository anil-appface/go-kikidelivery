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

run2:
	./"${APPLICATION}" 2

