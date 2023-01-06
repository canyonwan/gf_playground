package backend

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"backend/login" method:"POST"  tags:"登录" summary:"登录"`
	Account  string ` json:"account" v:"required#请输入账号|length:6,30#账号长度应当在:min到:max之间" dc:"账号"`
	Password string ` json:"password" v:"required#请输入密码|length:6,30#密码长度应当在:min到:max之间" dc:"密码"`
}

type LoginRes struct {
	Info interface{} `json:"info"`
}
