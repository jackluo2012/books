package access

import (
	"books/basic/redis"
	"fmt"
	r "github.com/go-redis/redis"
	"sync"
)

// # 负责定义、初始化等

var (
	s  *service
	ca *r.Client
	m  sync.RWMutex
)

// service 服务
type service struct {
}


// Service 用户服务类
type Service interface {
	// MakeAccessToken生成token
	MakeAccessToken(subject *Subject) (ret string, err error)

	// GetCacheAccessToken 获取缓存的token
	GetCacheAccessToken(subject *Subject) (ret string, err error)

	// DelUserAccessToken 清除用户token
	DelUserAccessToken(token string) (err error)

}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化用户服务层
func Init() {
	m.Lock()
	defer m.Unlock()
	if s != nil {
		return
	}
	ca = redis.GetRedis()
	s = &service{}
}
