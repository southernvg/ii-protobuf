# Make file for Protobuf
PROTOC=$(HOME)/bin/protoc

all: build
build:
	$(PROTOC) -I stackfinder stackfinder/stackfinder.proto --go_out=plugins=grpc:stackfinder

clean:
	rm -f stackfinder/stackfinder.pb.go


