module books/auth

go 1.13

require (
	books/basic v0.0.0-00010101000000-000000000000
	github.com/coreos/etcd v3.3.17+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/golang/protobuf v1.3.4
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
)

replace (
	books/basic => ../basic
	books/user-srv => ../user-srv
)
