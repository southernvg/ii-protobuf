# Make file for Protobuf
PROTOC=/usr/local/bin/protoc

all: build
build:
	$(PROTOC) -I stackfinder stackfinder/stackfinder.proto --go_out=plugins=grpc:stackfinder

clean:
	rm -f stackfinder/stackfinder.pb.go


