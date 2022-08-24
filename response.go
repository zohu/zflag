package zflag

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/zohu/zlog"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

var (
	EmptyData = make(map[string]string, 0)
	EmptyList = make([]map[string]string, 0)
)

func Done(c *gin.Context, res *ResponseBean) {
	defer func() {
		if r := recover(); r != nil {
			zlog.Error("panic", zap.Any("error", r))
			c.Abort()
		}
	}()
	if res.Kind == "" {
		ct := c.ContentType()
		switch ct {
		case "application/xml", "text/xml":
			res.Kind = ResponseKindXml
		case "application/json":
			res.Kind = ResponseKindJson
		default:
			res.Kind = ResponseKindJson
		}
	}
	switch res.Kind {
	case ResponseKindJson:
		c.JSON(http.StatusOK, res)
	case ResponseKindXml:
		c.XML(http.StatusOK, res)
	case ResponseKindString:
		c.String(http.StatusOK, fmt.Sprint(res.Data))
	default:
		c.Data(http.StatusOK, c.ContentType(), []byte(fmt.Sprint(res.Data)))
	}
	c.Abort()
}

func Success(data interface{}) *ResponseBean {
	return Response(1, data, "success", "", "")
}

func ResponseErr(flag Flag) *ResponseBean {
	return Response(flag.Code(), nil, flag.Message(), flag.Detail(), "")
}

func ParamErr(err error) *ResponseBean {
	e, ok := err.(validator.ValidationErrors)
	if !ok {
		if err.Error() == "EOF" {
			return ResponseErr(ErrParams.WithMessage("EOF"))
		}
		return ResponseErr(ErrParams.WithMessage(err.Error()))
	}
	var message []string
	for _, errc := range e {
		msg, err := ValiTrans.T(errc.Tag(), errc.Field(), errc.Param())
		if err == nil {
			message = append(message, msg)
		} else {
			message = append(message, errc.Error())
		}
	}
	return ResponseErr(ErrParams.WithMessage(strings.Join(message, ",")))
}

func Response(code int, data interface{}, message string, detail string, kind ResponseKind) *ResponseBean {
	if message == "" {
		message = "success"
	}
	return &ResponseBean{
		Kind:      kind,
		Version:   "1.0",
		Flag:      code,
		Data:      data,
		Message:   message,
		Detail:    detail,
		Timestamp: time.Now().UnixMilli(),
	}
}
