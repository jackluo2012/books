module books/payment-srv

go 1.13

require (
	books/basic v0.0.0-00010101000000-000000000000
	books/inventory-srv v0.0.0-00010101000000-000000000000
	books/orders-srv v0.0.0-00010101000000-000000000000
	books/plugins v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.3.4
	github.com/google/uuid v1.1.1
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.2.0
	github.com/micro/go-plugins/config/source/grpc v0.0.0-20200119172437-4fe21aa238fd
)

replace (
	books/basic => ../basic
	books/inventory-srv => ../inventory-srv
	books/orders-srv => ../orders-srv
	books/payment-srv => ../payment-srv
	books/plugins => ../plugins
)
