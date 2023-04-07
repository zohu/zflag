package zflag

type ResponseKind string

const (
	ResponseKindJson   ResponseKind = "json"
	ResponseKindXml    ResponseKind = "xml"
	ResponseKindString ResponseKind = "string"
	ResponseKindFile   ResponseKind = "file"
)

// ResponseBean
// @Description: 接口返回值定义
type ResponseBean struct {
	Kind    ResponseKind `json:"-" xml:"-"`
	Version string       `json:"-" xml:"version,attr"`
	Code    int          `json:"code"`
	Data    interface{}  `json:"data"`
	Message string       `json:"message"`
	Detail  string       `json:"detail"`
	T       string       `json:"t"`
}
