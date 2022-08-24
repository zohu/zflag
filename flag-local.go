package zflag

import "fmt"

type Flag interface {
	Code() int
	Message() string
	Detail() string
	String() string
	WithMessage(string) Flag
	WithDetail(string) Flag
}

type localFlag struct {
	flag    int
	message string
	detail  string
}

func (f *localFlag) Code() int {
	return f.flag
}

func (f *localFlag) Message() string {
	return f.message
}

func (f *localFlag) Detail() string {
	return f.detail
}

func (f *localFlag) WithMessage(msg string) Flag {
	lf := &localFlag{
		flag:    f.flag,
		message: f.message,
		detail:  f.detail,
	}
	if lf.message != "" {
		lf.message = fmt.Sprintf("%s %s", lf.message, msg)
	} else {
		lf.message = msg
	}
	return lf
}

func (f *localFlag) WithDetail(dt string) Flag {
	lf := &localFlag{
		flag:    f.flag,
		message: f.message,
		detail:  f.detail,
	}
	if lf.detail != "" {
		lf.detail = fmt.Sprintf("%s %s", lf.detail, dt)
	} else {
		lf.detail = dt
	}
	return lf
}

func (f *localFlag) String() string {
	if f.detail != "" {
		return fmt.Sprintf("%d:%s %s", f.flag, f.message, f.detail)
	}
	return fmt.Sprintf("%d:%s", f.flag, f.message)
}
