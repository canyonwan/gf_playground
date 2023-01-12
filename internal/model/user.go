package model

// UserInfoGetOutput 获取用户信息
type UserInfoGetOutput struct {
	Id           uint     `json:"id" dc:"用户ID"`
	Account      string   `json:"account" dc:"帐号"`
	Username     string   `json:"username" dc:"用户名"`
	Gender       int8     `json:"gender" dc:"性别"`
	Avatar       string   `json:"avatar" dc:"头像"`
	Email        string   `json:"email" dc:"邮箱"`
	Phone        string   `json:"phone" dc:"手机号"`
	RoleIds      string   `json:"roleIds" dc:"角色Ids"`
	Permissions  []string `json:"permissions" dc:"权限列表"`
	IsSuperAdmin int8     `json:"isSuperAdmin" dc:"是否超级管理员"`
}
