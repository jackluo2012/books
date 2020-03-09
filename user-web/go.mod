module books/user-web

go 1.13

replace (
	books/auth => ../auth
	books/basic => ../basic
	books/plugins => ../plugins
	books/user-srv => ../user-srv
	google.golang.org/grpc v1.27.1 => google.golang.org/grpc v1.26.0
)

require (
	books/auth v0.0.0-00010101000000-000000000000
	books/basic v0.0.0-00010101000000-000000000000
	books/plugins v0.0.0-00010101000000-000000000000
	books/user-srv v0.0.0-00010101000000-000000000000
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.2.0
	github.com/micro/go-plugins/config/source/grpc/v2 v2.0.3
	github.com/micro/go-plugins/wrapper/breaker/hystrix/v2 v2.0.3
)
