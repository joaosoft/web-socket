server:
	go run ./examples/main.go -mode server

client:
	go run ./examples/main.go -mode client

build:
	go build .

fmt:
	go fmt ./...

vet:
	go vet ./*

gometalinter:
	gometalinter ./*