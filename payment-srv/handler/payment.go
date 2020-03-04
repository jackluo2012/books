package handler

import (
	"books/payment-srv/model/payment"
	"context"

	"github.com/micro/go-micro/util/log"
	proto "books/payment-srv/proto/payment"
)


var(
	paymentService payment.Service
)

type Payment struct{}

// Init 初始化handler
func Init() {
	paymentService, _ = payment.GetService()
}

// New 新增订单
func (e *Payment) PayOrder(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	log.Log("[PayOrder] 收到支付请求")
	err = paymentService.PayOrder(req.OrderId)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error(),
		}
		return
	}

	rsp.Success = true
	return
}
