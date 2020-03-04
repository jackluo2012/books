module books/inventory-srv

go 1.13

replace books/basic => ../basic

require (
	books/basic v0.0.0-00010101000000-000000000000
	github.com/go-log/log v0.1.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.18.0
)
