package respstatus

import "fmt"

const (
	// System Code
	SYSTEM_ERROR_NO = 0
)

type ErrNo struct {
	ErrCode uint32
	ErrMsg  string
}

// Error 实现 error 类型
func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code uint32, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success    = NewErrNo(SYSTEM_ERROR_NO, "Success")
	ServiceErr = NewErrNo(SYSTEM_ERROR_NO+1, "Service is unable to start successfully")
	ParamErr   = NewErrNo(SYSTEM_ERROR_NO+2, "Wrong Parameter has been given")
)
