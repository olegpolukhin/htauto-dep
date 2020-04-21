# Compiles protos
protos: 
	@echo "Building protos ..."
	protoc -I service/ service/*.proto --go_out=plugins=grpc:proto
