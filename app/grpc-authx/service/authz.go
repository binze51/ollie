package service

import (
	"strings"

	"k8s.io/klog"
)

// batchCheck 用户 对api方法的role权限校验，支持多role
func (c *ServiceImpl) batchCheck(roleIds, act, obj string) bool {
	var checkReq [][]any
	for _, v := range strings.Split(roleIds, ",") {
		checkReq = append(checkReq, []any{v, obj, act})
	}

	result, err := c.enforcer.BatchEnforce(checkReq)
	if err != nil {
		klog.Error("Casbin enforce error", err.Error())
		return false
	}

	for _, v := range result {
		if v {
			return true
		}
	}

	return false
}
