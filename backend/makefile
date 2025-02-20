all: run

run:
	go run main.go

build:
	go build -v hestia/jobfair/api
test:
	go test -v ./...

proto-server:
	protoc --proto_path=../proto --go_out=. --go_opt=module=hestia/jobfair/api --go-grpc_out=. --go-grpc_opt=module=hestia/jobfair/api ../proto/*.proto

clean:
	git clean -Xfd