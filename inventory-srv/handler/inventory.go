package handler

import (
	"books/inventory-srv/model/inventory"
	proto "books/inventory-srv/proto/inventory"
	"context"
	"github.com/go-log/log"
)

var (
	invService inventory.Service
)

type Inventory struct{}

// Init 初始化handler
func Init() {
	invService, _ = inventory.GetService()
}

// Sell 库存销存
func (e *Inventory) Sell(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	id, err := invService.Sell(req.BookId, req.UserId)
	if err != nil {
		log.Logf("[Sell] 销存失败，bookId：%d，userId: %d，%s", req.BookId, req.UserId, err)
		rsp.Success = false
		return
	}

	rsp.InvH = &proto.InvHistory{Id: id,}
	rsp.Success = true
	return nil
}

// Confirm 库存销存 确认
func (e *Inventory) Confirm(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	err = invService.Confirm(req.HistoryId, int(req.HistoryState))
	if err != nil {
		log.Logf("[Confirm] 确认销存失败，%s", err)
		rsp.Success = false
		return
	}
	rsp.Success = true
	return nil
}
