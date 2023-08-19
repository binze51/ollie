package event

import (
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/encoders/protobuf"
	"github.com/spf13/viper"
)

var (
	eccli  *nats.EncodedConn
	onceEc sync.Once
)

func InitProtoEncodedConn() (ec *nats.EncodedConn, err error) {
	onceEc.Do(func() {
		// url := fmt.Sprintf("nats://127.0.0.1:%d", port)
		nc, err := nats.Connect(viper.GetString("nats.url"))
		if err != nil {
			return
		}

		// protobuf 序列化
		ec, err = nats.NewEncodedConn(nc, protobuf.PROTOBUF_ENCODER)
		if err != nil {
			return
		}
		eccli = ec
	})

	return eccli, err
}
