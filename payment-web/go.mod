module books/payment-web

go 1.13

require (
	books/auth v0.0.0-00010101000000-000000000000
	books/basic v0.0.0-00010101000000-000000000000
	books/payment-srv v0.0.0-00010101000000-000000000000
	books/plugins v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.2.0
)

replace (
	books/auth => ../auth
	books/basic => ../basic
	books/inventory-srv => ../inventory-srv
	books/orders-srv => ../orders-srv
	books/payment-srv => ../payment-srv
	books/plugins => ../plugins
)
