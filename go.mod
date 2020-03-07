module books

go 1.13

require (
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/config/source/grpc/v2 v2.0.3
	golang.org/x/net v0.0.0-20200226051749-491c5fce7268 // indirect
	golang.org/x/sys v0.0.0-20200223170610-d5e6a3e2c0ae // indirect
	google.golang.org/genproto v0.0.0-20200225123651-fc8f55426688 // indirect
	google.golang.org/grpc v1.27.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace google.golang.org/grpc v1.27.1 => google.golang.org/grpc v1.26.0
