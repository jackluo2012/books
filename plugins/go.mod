module books/plugins

go 1.13

require (
	books/basic v0.0.0-00010101000000-000000000000
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.1.1
	github.com/gorilla/sessions v1.2.0
	github.com/micro/go-micro v1.18.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-client-go v2.22.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	go.uber.org/zap v1.14.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace (
	books/basic => ../basic
	books/inventory-srv => ../inventory-srv
	books/orders-srv => ../orders-srv
	books/payment-srv => ../payment-srv
)
