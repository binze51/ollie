package config

import (
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// InitConfig - read configfile and ENV variables
func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	if _, err := os.Stat("./biz"); err == nil {
		viper.AddConfigPath("./biz")
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// hot reloading
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		klog.Info("config file changed:", e.Name)
	})
}
