.PHONY: all

all: clean deps genProtoBuf rungRPCServer
	 
clean:
	@-rm -f ./imdbprotopb/imdb_service_grpc.pb.go ./imdbprotopb/imdb_service.pb.go 
	@echo "[OK] removed all old protobuf files!"

.PHONY: deps

deps:
	go mod tidy

.PHONY: genProtoBuf

genProtoBuf:
	protoc --python_out=imdbprotopb --go_out=imdbprotopb --go-grpc_out=imdbprotopb imdbproto/*.proto
	@cp imdbprotopb/github.com/dhananjayksharma/dkgosqlimdbapi/imdbproto/* imdbprotopb/.
rungRPCServer:
	go run imdbgrpc-server/imdbgrpc-server.go