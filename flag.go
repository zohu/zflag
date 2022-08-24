package zflag

func NewFlag(flag int, message string, detail string) Flag {
	if flag < 1000 {
		panic("flag < 1000 is not allowed, please use defaultFlag 'zflag.xxx'")
	}
	return &localFlag{
		flag:    flag,
		message: message,
		detail:  detail,
	}
}

func defaultFlag(flag int, message string, detail string) Flag {
	return &localFlag{
		flag:    flag,
		message: message,
		detail:  detail,
	}
}
