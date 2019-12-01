package cerror

var successCode = ErrCode{0, 200}

type ErrCode struct {
	c uint
	s int
}

func NewErrCode(code uint, httpStatusCode int) ErrCode {
	return ErrCode{
		c: code,
		s: httpStatusCode,
	}
}
