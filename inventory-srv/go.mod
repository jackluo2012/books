module books/inventory-srv

go 1.13

replace (
	books/basic => ../basic
	books/plugins => ../plugins
)

require (
	books/basic v0.0.0-00010101000000-000000000000
	books/plugins v0.0.0-00010101000000-000000000000
	github.com/go-log/log v0.1.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.3.2
	github.com/micro/cli v0.2.0
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.2.0
	github.com/micro/go-plugins/config/source/grpc/v2 v2.0.3
)
