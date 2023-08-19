package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"ollie/pkg/authz"
	"ollie/pkg/db"

	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"google.golang.org/grpc/health/grpc_health_v1"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

func InitServiceContainer(ctx context.Context) (*ServiceImpl, error) {
	db, err := db.InitDB(viper.GetString("db.pgsqlsdn"))
	if err != nil {
		return nil, err
	}
	enforcer, err := authz.InitEnforcer(db)
	if err != nil {
		return nil, authxErrErr
	}
	// if viper.GetBool("gorm.enableAutoMigrate") {
	// 	err = db.AutoMigrate(new(authx.User))
	// 	if err != nil {
	// 		return nil, authxErrErr
	// 	}
	// }

	return &ServiceImpl{
		db:       db,
		enforcer: enforcer,
	}, nil
}

type ServiceImpl struct {
	grpc_health_v1.UnimplementedHealthServer
	db *gorm.DB
	// api权限验证器
	enforcer *casbin.SyncedEnforcer

	// account rpc
	// accountSvc
}

// Close 回收所有依赖sdk的tcp连接
func (c *ServiceImpl) Close() {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	var errs []error
	go func() {
		defer wg.Done()
		s, _ := c.db.DB()
		if err := s.Close(); err != nil {
			errs = append(errs, err)
		}
		if viper.GetBool("casbin.autoLoad") {
			c.enforcer.StopAutoLoadPolicy()
		}
	}()

	wg.Wait()

	var closeErr error
	for _, err := range errs {
		if closeErr == nil {
			closeErr = err
		} else {
			closeErr = fmt.Errorf("%v | %v", closeErr, err)
		}
	}

	if closeErr != nil {
		klog.Error(closeErr.Error())
	}
}
