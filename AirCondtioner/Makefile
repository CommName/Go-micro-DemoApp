
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/AirCondtioner/AirCondtioner.proto

.PHONY: build
build: proto

	go build -o AirCondtioner-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t AirCondtioner-srv:latest
