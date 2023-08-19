package authz

import (
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	onceEnforce sync.Once
	enforcer    *casbin.SyncedEnforcer
)

func InitEnforcer(db *gorm.DB) (e *casbin.SyncedEnforcer, err error) {
	onceEnforce.Do(func() {
		e, err = casbin.NewSyncedEnforcer(viper.GetString("casbin.modelFile"))
		if err != nil {
			return
		}
		e.EnableLog(viper.GetBool("casbin.debug"))

		adapter, err := gormadapter.NewAdapterByDBUseTableName(db, "all", "authx_rules")
		if err != nil {
			return
		}

		err = e.InitWithModelAndAdapter(e.GetModel(), adapter)
		if err != nil {
			return
		}
		enforcer = e
	})
	return enforcer, err
}
