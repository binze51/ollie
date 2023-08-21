package respstatus

import (
	"errors"

	"ollie/kitex_gen/common/respstatus"
)

// BuildStatusResp build StatusResp from error
func BuildStatusResp(err error) *respstatus.StatusResp {
	if err == nil {
		return baseResp(Success)
	}

	e := ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err ErrNo) *respstatus.StatusResp {
	return &respstatus.StatusResp{Code: uint32(err.ErrCode), Msg: err.ErrMsg}
}
