module books/plugins

go 1.13

require (
	books/basic v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/google/uuid v1.1.1
	github.com/gorilla/sessions v1.2.0
	github.com/micro/go-micro v1.18.0
)

replace (
	books/basic => ../basic
	books/inventory-srv => ../inventory-srv
	books/orders-srv => ../orders-srv
	books/payment-srv => ../payment-srv
)
