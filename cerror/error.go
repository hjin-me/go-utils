package cerror

type Error struct {
	Code ErrCode `json:"err_code"`
	Msg  string  `json:"err_msg"`
	err  error
}

func (e *Error) Unwrap() error {
	return e.err
}
func (e Error) Error() string {
	return e.Msg
}
func (e *Error) String() string {
	return e.Msg
}
func (e *Error) StatusCode() int {
	return e.Code.s
}
func New(code ErrCode, msg string, err error) Error {
	return Error{code, msg, err}
}
