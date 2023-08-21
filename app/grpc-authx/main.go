// main incoming
package main

import (
	"context"
	"net"

	"ollie/app/grpc-authx/service"
	"ollie/kitex_gen/authx/authxservice"
	"ollie/pkg/config"
	"ollie/pkg/logger"
	"ollie/pkg/shutdown"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/spf13/viper"
)

func init() {
	config.InitConfig()
	logger.InitLogger()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	serviceContainer, err := service.InitServiceContainer(ctx)
	if err != nil {
		panic(err)
	}
	defer shutdown.GracefulStop(cancel, serviceContainer.Close)

	// otel provider
	// p := provider.NewOpenTelemetryProvider(
	// 	provider.WithServiceName(viper.GetString("app.name")),
	// 	provider.WithExportEndpoint(viper.GetString("otel.endpoint")),
	// 	provider.WithInsecure(),
	// )
	// defer p.Shutdown(context.Background())

	opts := kitexOpts()
	svr := authxservice.NewServer(serviceContainer, opts...)

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexOpts() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", ":"+viper.GetString("app.port"))
	if err != nil {
		panic(err)
	}
	// server side middleware
	opts = append(opts, server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxConnections: viper.GetInt("app.limit.maxConnects"), MaxQPS: viper.GetInt("app.limit.maxQPS")}),
		// server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: viper.GetString("app.name")}),
	)
	return
}
