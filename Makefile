# Compiles protos
protos: 
	@echo "Building protos ..."
	#protoc -I service/ service/*.proto --go_out=plugins=grpc:service
	#protoc --proto_path=third_party --go_out=plugins=grpc:proto service.proto
	protoc -I service/ service/*.proto --go_out=plugins=grpc:proto