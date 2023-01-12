# dkgosqlimdbapi
Folder:
dkgosqlimdbapi/
	imdbapi-service
	imdbgrpc-client
	imdbgrpc-server
	imdbproto
	imdbprotopb
	
1. Setup MySQL database
	import database from databases
	
2. Run make in folder: dkgosqlimdbapi
	make genProtoBuf
	
3 Setup mysql password into 
	imdbgrpc-server/config-local.yml
	
	and
	
	imdbapi-service/config-local.yml
	
4. Run make into imdbapi-service (REST-API SERVER)
	make
	
5. Run make into imdbgrpc-server (gRPC SERVER)
	make

6. Run one by one functiona into folder imdbgrpc-client
	addPerson(gRPCConn) // Unary
	searchPerson(gRPCConn) // Server Streaming
	addManyPerson(gRPCConn) // Client Streaming
	searchPersonById(gRPCConn) // Bidirectional Server Streaming
