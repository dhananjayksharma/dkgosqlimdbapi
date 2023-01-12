module github.com/dhananjayksharma/dkgosqlimdbapi/imdbgrpc-client

go 1.18

replace github.com/dhananjayksharma/dkgosqlimdbapi/imdbprotopb v1.0.0 => ./../imdbprotopb

require google.golang.org/protobuf v1.28.1 // indirect

require (
	github.com/dhananjayksharma/dkgosqlimdbapi/imdbprotopb v1.0.0
	google.golang.org/grpc v1.51.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20230106154932-a12b697841d9 // indirect
)
