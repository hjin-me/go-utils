package cerror

import "strconv"

var successCode = ErrCode{0, 200}
var internalErrCode uint = 500

type ErrCode struct {
	// 错误码
	c uint
	// http status code
	s int
}

func (e ErrCode) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(e.c), 10)), nil
}
func (e ErrCode) Code() uint {
	return e.c
}

func NewErrCode(code uint, httpStatusCode int) ErrCode {
	return ErrCode{
		c: code,
		s: httpStatusCode,
	}
}
