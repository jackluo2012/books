package handler

import (
	"books/auth/model/access"
	auth "books/auth/proto/auth"
	z "books/plugins/zap"
	"context"
	"strconv"
)

var (
	accessService access.Service
	log           = z.GetLogger()
)

//Init 初始化handler
func Init() {
	var err error
	accessService, err = access.GetService()

	if err != nil {
		log.Fatal("[Init] 初始化Handler错误," +err.Error())
		return
	}
}

type Service struct{}

// MakeAccessToken 生成token
func (s *Service) MakeAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Info("[MakeAccessToken] 收到创建token请求")

	token, err := accessService.MakeAccessToken(&access.Subject{
		ID:   strconv.FormatUint(req.UserId, 10),
		Name: req.UserName,
	})
	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Debug("[MakeAccessToken] token生成失败，err：%s" + err.Error())
		return err
	}

	rsp.Token = token
	return nil
}

// DelUserAccessToken 清除用户token
func (s *Service) DelUserAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Info("[DelUserAccessToken] 清除用户token")
	err := accessService.DelUserAccessToken(req.Token)
	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Info("[DelUserAccessToken] 清除用户token失败，err" + err.Error())
		return err
	}

	return nil
}

// GetCachedAccessToken 获取缓存的token
func (s *Service) GetCachedAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Info("[GetCachedAccessToken] 获取缓存的token " + strconv.FormatUint(req.UserId, 10))
	token, err := accessService.GetCacheAccessToken(&access.Subject{
		ID: strconv.FormatInt(int64(req.UserId), 10),
	})
	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Info("[GetCachedAccessToken] 获取缓存的token失败，err:" + err.Error())
		return err
	}

	rsp.Token = token
	return nil
}
