package a_user

import (
	"github.com/gogf/gf-demos/app/service/user"
	"github.com/gogf/gf-demos/library/response"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/util/gvalid"
)

// 用户API管理对象
type Controller struct { }

// 用户注册接口
func (c *Controller) SignUp(r *ghttp.Request) {
	if err := s_user.SignUp(r.GetPostMap()); err != nil {
		response.Json(r, 1, err.Error())
	} else {
		response.Json(r, 0, "ok")
	}
}

// 用户登录接口
func (c *Controller) SignIn(r *ghttp.Request) {
	data  := r.GetPostMap()
	rules := map[string]string {
		"passport"  : "required",
		"password"  : "required",
	}
	msgs  := map[string]interface{} {
		"passport" : "账号不能为空",
		"password" : "密码不能为空",
	}
	if e := gvalid.CheckMap(data, rules, msgs); e != nil {
		response.Json(r, 1, e.String())
	}
	if err := s_user.SignIn(data["passport"], data["password"], r.Session); err != nil {
		response.Json(r, 1, err.Error())
	} else {
		response.Json(r, 0, "ok")
	}
}

// 判断用户是否已经登录
func (c *Controller) IsSignedIn(r *ghttp.Request) {
	if s_user.IsSignedIn(r.Session) {
		response.Json(r, 0, "ok")
	} else {
		response.Json(r, 1, "")
	}
}

// 用户注销/退出接口
func (c *Controller) SignOut(r *ghttp.Request) {
	s_user.SignOut(r.Session)
	response.Json(r, 0, "ok")
}

// 检测用户账号接口(唯一性校验)
func (c *Controller) CheckPassport(r *ghttp.Request) {
	passport := r.Get("passport")
	if e := gvalid.Check(passport, "required", "请输入账号"); e != nil {
		response.Json(r, 1, e.String())
	}
	if s_user.CheckPassport(passport) {
		response.Json(r, 0, "ok")
	}
	response.Json(r, 1, "账号已经存在")
}

// 检测用户昵称接口(唯一性校验)
func (c *Controller) CheckNickName(r *ghttp.Request) {
	nickname := r.Get("nickname")
	if e := gvalid.Check(nickname, "required", "请输入昵称"); e != nil {
		response.Json(r, 1, e.String())
	}
	if s_user.CheckNickName(r.Get("nickname")) {
		response.Json(r, 0, "ok")
	}
	response.Json(r, 1, "昵称已经存在")
}