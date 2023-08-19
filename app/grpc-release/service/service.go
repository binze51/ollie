// Package service ...s
package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"ollie/kitex_gen/release"
	"ollie/pkg/authz"
	"ollie/pkg/db"

	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"google.golang.org/grpc/health/grpc_health_v1"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

// InitServiceContainer 初始化服务容器
func InitServiceContainer(ctx context.Context) (*ServiceImpl, error) {
	db, err := db.InitDB(viper.GetString("db.pgsqlsdn"))
	if err != nil {
		return nil, err
	}
	enforcer, err := authz.InitEnforcer(db)
	if err != nil {
		return nil, ReleaseErr
	}
	if viper.GetBool("gorm.enableAutoMigrate") {
		err = db.AutoMigrate(
			new(release.X_Release), new(release.X_CIBuildService), new(release.X_ReleaseWebOpLog))
		if err != nil {
			return nil, ReleaseErr
		}
	}

	return &ServiceImpl{
		db:       db,
		Ticker:   time.NewTicker(time.Duration(viper.GetInt64("app.tickDuration")) * time.Second),
		enforcer: enforcer,
	}, nil
}

type ServiceImpl struct {
	grpc_health_v1.UnimplementedHealthServer
	db *gorm.DB
	// api权限验证器
	enforcer *casbin.SyncedEnforcer

	// ticker 定时器
	Ticker *time.Ticker
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

		c.Ticker.Stop()
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
