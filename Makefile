GOPATH=$(shell go env GOPATH)

gen-proto:
	protoc -I. \
      -I/usr/local/include \
      -I${GOPATH}/src/ \
      --go_out=plugins=grpc:../../../ \
      ./proto/**/*.proto