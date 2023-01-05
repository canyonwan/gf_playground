package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ContextUser 请求上下文中的用户信息
type ContextUser struct {
	Id           uint
	Account      string
	UserName     string
	Avatar       string
	IsSuperAdmin uint8
}

// Context 请求上下文结构
type Context struct {
	Session *ghttp.Session // 当前Session管理对象
	User    *ContextUser   // 上下文用户信息
	Data    g.Map          // 自定义KV变量,不固定
}
