PROTOCHEAD=protoc -I/usr/local/include -I. -I${GOPATH}/src -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

all: proto gw swagger

proto:
	@$(PROTOCHEAD) --go_out=plugins=grpc:. ./protos/*.proto

gw:
	@$(PROTOCHEAD) --grpc-gateway_out=logtostderr=true:. ./protos/*.proto

swagger:
	@$(PROTOCHEAD) --swagger_out=logtostderr=true:. ./protos/*.proto