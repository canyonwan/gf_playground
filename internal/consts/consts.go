package consts

const (
	ContextKey               = "ContextKey" // 上下文变量存储键名
	FileMaxUploadCountMinute = 10           // 同一用户1分钟之内最大上传数量
	AuthFailMsg              = "登录超时或认证失败"
	LoginPath                = "/backend/login" // 登录路由
	LogoutPath               = "/backend/logout"
	GTokenAdminPrefix        = "Admin:" // gtoken登录 管理后台 前缀区分
	CtxAdminId               = "CtxAdminId"
	CtxAdminName             = "CtxAdminName"
	CtxAdminIsAdmin          = "CtxAdminIsAdmin"
	CtxAdminRoleIds          = "CtxAdminRoleIds"
)
