package zflag

// default flag -1 - 999
var (
	Ok     = defaultFlag(1, "", "")
	ErrNil = defaultFlag(-1, "未知错误", "")

	ErrParams        = defaultFlag(900, "param error", "")
	ErrVerifyAuth    = defaultFlag(901, "用户异常", "")
	ErrLogin         = defaultFlag(902, "登录失败", "")
	ErrNoPermissions = defaultFlag(903, "无权访问", "请联系管理员赋权")
	ErrVersionLow    = defaultFlag(904, "版本号过低", "请下载最新版本")
	ErrSign          = defaultFlag(905, "签名错误", "请检查签名")
	ErrXPwd          = defaultFlag(906, "鉴权失败", "")
	ErrAuth          = defaultFlag(907, "登录状态已过期", "请重新登录再试")
)
