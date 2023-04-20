package zflag

// default flag -1 - 999
var (
	Ok     = defaultFlag(1, "", "")
	ErrNil = defaultFlag(-1, "未知错误", "")

	ErrParams        = defaultFlag(9000, "param error", "")
	ErrVerifyAuth    = defaultFlag(9001, "用户异常", "")
	ErrLogin         = defaultFlag(9002, "登录失败", "")
	ErrNoPermissions = defaultFlag(9003, "无权访问", "请联系管理员赋权")
	ErrVersionLow    = defaultFlag(9004, "版本号过低", "请下载最新版本")
	ErrSign          = defaultFlag(9005, "签名错误", "请检查签名")
	ErrXPwd          = defaultFlag(9006, "鉴权失败", "")
	ErrAuth          = defaultFlag(9007, "登录状态已过期", "请重新登录再试")
)
