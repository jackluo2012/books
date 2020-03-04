package handler

import (
	us "books/user-srv/model/user"
	s "books/user-srv/proto/user"
	"context"
	"log"
)

type User struct{}

var (
	userService us.Service
)

// Init 初始化handler

func Init() {
	var err error
	userService, err = us.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

func (e *User) QueryUserByName(ctx context.Context, req *s.Request, resp *s.Response) error {
	//panic("implement me")
	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		resp.Success = false
		resp.Error = &s.Error{
			Code:   0,
			Detail: "",
		}
		return nil
	}
	resp.User = user
	resp.Success = true
	return nil
}
