module books/user-web

go 1.13

require (
	books/auth v0.0.0-00010101000000-000000000000
	books/plugins v0.0.0-00010101000000-000000000000
	books/user-srv v0.0.0-00010101000000-000000000000
	github.com/go-log/log v0.1.0
	github.com/micro/go-micro v1.18.0
)

//replace github.com/books/user-srv => /Users/jackluo/works/go/src/books/user-srv

replace (
	books/auth => ../auth
	books/basic => ../basic
	books/plugins => ../plugins
	books/user-srv => ../user-srv
)
