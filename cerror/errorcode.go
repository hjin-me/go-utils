package cerror

var successCode = ErrCode{0, 200}

type ErrCode struct {
	// 错误码
	c uint
	// http status code
	s int
}

func NewErrCode(code uint, httpStatusCode int) ErrCode {
	return ErrCode{
		c: code,
		s: httpStatusCode,
	}
}
