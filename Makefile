.PHONY: compile
PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go

# If $GOPATH/bin/protoc-gen-go does not exist, we'll run this command to install
# it.
$(PROTOC_GEN_GO):
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -d github.com/envoyproxy/protoc-gen-validate
	go get -u github.com/googleapis/googleapis


proto:
	echo "Build Proto"
	protoc -I. -I /usr/local/include -I${GOPATH}/src -I${GOPATH}/src/hangman/protos -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis  -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		protos/hangman_api/v1/hangman_api.proto --go_out=plugins=grpc:. --validate_out="lang=go:${GOPATH}/src"
	protoc -I. -I /usr/local/include -I${GOPATH}/src -I${GOPATH}/src/hangman/protos -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		protos/hangman_api/v1/hangman_api.proto --grpc-gateway_out=logtostderr=true:.
	protoc -I. -I /usr/local/include -I${GOPATH}/src -I${GOPATH}/src/hangman/protos -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		protos/hangman_api/v1/hangman_api.proto --swagger_out=logtostderr=true:.
	swagger-codegen generate -i ${GOPATH}/src/hangman/protos/hangman_api/v1/hangman_api.swagger.json -l html -o ${GOPATH}/src/hangman/protos/hangman_api/v1/doc

build:
	echo "Build Hangman"
	go build -o bin/main main.go

run:
	go run main.go


#compile:
#	echo "Compiling for every OS and Platform"
#	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
#	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
#	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

all: proto build