package settings

import (
	
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name    string `mapstructure:"name"`
		Mode    string `mapstructure:"mode"`
		Port    int    `mapstructure:"port"`
		Version string `mapstructure:"version"`
		StartTime  string `mapstructure:"start_time"`
		MachineID int64 `mapstructure:"machine_id"`
	}
	Log struct {
		// `mapstructure:"xxxx"`确保正确映射
		Level      string `mapstructure:"level"`
		Filename   string `mapstructure:"filename"`
		MaxSize    int    `mapstructure:"max_size"`
		MaxAge     int    `mapstructure:"max_age"`
		MaxBackups int    `mapstructure:"max_backups"`
	}
	Mysql struct {
		Host         string `mapstructure:"host"`
		Port         string `mapstructure:"port"`
		User         string `mapstructure:"user"`
		Password     string `mapstructure:"password"`
		Name         string `mapstructure:"name"`
		MaxIdleConns int    `mapstructure:"max_idle_conns"`
		MaxOpenConns int    `mapstructure:"max_open_conns"`
	}
	Redis struct {
		Host     string `mapstructure:"host"`
		Password string `mapstructure:"password"`
		Port     int    `mapstructure:"port"`
		DB       int    `mapstructure:"db"`
		PoolSize int    `mapstructure:"pool_size"`
	}
}

var AppConfig *Config

func Init() (err error) {
	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath(".")
	viper.SetConfigFile("./conf/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		// zap.L().Fatal("Error reading config file: %v", zap.Error(err))
		fmt.Printf("viper.ReadInConfig() failed,err: %v\n", err)
		return err
	}
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		fmt.Printf("viper.Unmarshal failed,err:%v\n", err)
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件被修改了...")
		if err := viper.Unmarshal(&AppConfig); err != nil {
			fmt.Printf("viper.Unmarshal failed,err:%v\n", err)
			return
		}
	})

	
	return nil
}
