package utils

import (
	"time"

	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func init() {
	sf = sonyflake.NewSonyflake(sonyflake.Settings{
		// id时间位起始时间
		StartTime: time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
	})
	if sf == nil {
		panic("sonyflake not created")
	}
}

// Get unique id from Twitter's Snowflake
func MustID() uint64 {
	id, err := sf.NextID()
	if err == nil {
		return id
	}

	sleep := 1
	for {
		// 毫秒匹配sonyflake算法
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		id, err := sf.NextID()
		if err == nil {
			return id
		}
		// 指数退避 重试
		sleep *= 2
	}
}
