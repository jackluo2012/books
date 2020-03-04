package model

import "books/auth/model/access"

// # 业务模型初始化入口
// Init 初始化模型层
func Init() {
	access.Init()
}
