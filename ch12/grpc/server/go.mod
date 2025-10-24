module github.com/wrelin/grpc/server

go 1.25.1

replace github.com/wrelin/proto => ../../proto

require (
	github.com/wrelin/proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.76.0
)

require (
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	google.golang.org/genproto v0.0.0-20210701191553-46259e63a0a9 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)
