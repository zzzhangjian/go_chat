package router

import (
	"github.com/gogf/gf/g"
	"go_chat/app/api/user"
)

func init() {
	// 用户模块 路由注册 - 使用执行对象注册方式
	g.Server().BindObject("/user", new(a_user.Controller))
}