package utils

import (
	"sync"
	"time"
)

// 带过期时间的sync.map
type ExpMap struct {
	Map sync.Map
}

// sync.Map里的值结构要实现该方法
type CheckExp interface {
	IsTimeOut() bool
}

//	func (v *va) IsTimeOut() bool {
//		return time.Since(v.LoginTime) > v.ExpiredTime
//	}

func NewExpMap(checkTime time.Duration) *ExpMap {
	expMap := ExpMap{}
	go func() {
		time.Sleep(checkTime)
		expMap.Map.Range(func(key, value any) bool {
			v := value.(CheckExp)
			if v.IsTimeOut() {
				expMap.Map.Delete(key)
			}
			return true
		})
	}()
	return &expMap
}
