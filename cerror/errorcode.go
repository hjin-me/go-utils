package cerror

import "strconv"

var successCode = ErrCode{0, 200}

type ErrCode struct {
	// 错误码
	c uint
	// http status code
	s int
}

func (e ErrCode) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(e.c), 10)), nil
}

func NewErrCode(code uint, httpStatusCode int) ErrCode {
	return ErrCode{
		c: code,
		s: httpStatusCode,
	}
}
