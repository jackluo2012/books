package payment

import (
	"books/basic/common"
	invS "books/inventory-srv/proto/inventory"
	ordS "books/orders-srv/proto/order"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"sync"
)

var (
	s            *service
	invClient    invS.InventoryService
	ordSClient   ordS.OrdersService
	m            sync.RWMutex
	payPublisher micro.Publisher
)

// service 服务
type service struct {
}

// Service 服务类
type Service interface {
	// PayOrder 支付订单
	PayOrder(orderId int64) (err error)
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化库存服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	invClient = invS.NewInventoryService("mu.micro.book.srv.inventory", client.DefaultClient)
	ordSClient = ordS.NewOrdersService("mu.micro.book.srv.order", client.DefaultClient)
	payPublisher = micro.NewPublisher(common.TopicPaymentDone, client.DefaultClient)
	s = &service{}
}
