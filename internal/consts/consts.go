package consts

const (
	ContextKey               = "ContextKey" // 上下文变量存储键名
	FileMaxUploadCountMinute = 10           // 同一用户1分钟之内最大上传数量
	AuthFailMsg              = "登录超时或认证失败"
	AdminLoginPath           = "/login" // 管理端登录路由
	LoginPath                = "/login" // app端登录路由
	AdminLogoutPath          = "/logout"
	LogoutPath               = "/logout"
	GTokenAdminPrefix        = "Admin:" // gtoken登录 管理后台 前缀区分
	CtxAdminId               = "CtxAdminId"
	CtxAdminAccount          = "CtxAdminAccount"
	CtxAdminIsSuperAdmin     = "CtxAdminIsSuperAdmin"
	CtxAdminRoleIds          = "CtxAdminRoleIds"
)
