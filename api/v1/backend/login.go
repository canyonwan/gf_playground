package backend

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" method:"POST"  tags:"登录" summary:"登录"`
	Account  string ` json:"account" v:"required#请输入账号|length:6,30#账号长度应当在:min到:max之间" dc:"账号"`
	Password string ` json:"password" v:"required#请输入密码|length:6,30#密码长度应当在:min到:max之间" dc:"密码"`
}

type LoginRes struct {
	Type         string `json:"type"`
	Token        string `json:"token"`
	ExpireIn     int    `json:"expireIn"`
	IsSuperAdmin int    `json:"isSuperAdmin"` //是否超管
	RoleIds      string `json:"roleIds"`      //角色
	//Permissions []entity.PermissionInfo `json:"permissions"` //权限列表
}
