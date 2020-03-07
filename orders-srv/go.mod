module books/orders-srv

go 1.13

require (
	books/basic v0.0.0-00010101000000-000000000000
	books/inventory-srv v0.0.0-00010101000000-000000000000
	books/payment-srv v0.0.0-00010101000000-000000000000
	books/plugins v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.3.4
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.2.0
	github.com/micro/go-plugins/config/source/grpc/v2 v2.0.3
)

replace (
	books/basic => ../basic
	books/inventory-srv => ../inventory-srv
	books/orders-srv => ../orders-srv
	books/payment-srv => ../payment-srv
	books/plugins => ../plugins
)
