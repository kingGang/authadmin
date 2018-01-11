package models

type JRResultCode int

const (
	JRCodeSucc JRResultCode = iota
	JRCodeFailed
	JRCode302 = 302
	JRCode401 = 401
)

type JsonResult struct {
	Code JRResultCode `json:"code"`
	Msg  string       `json:"msg"`
	Obj  interface{}  `json:"obj"`
}
