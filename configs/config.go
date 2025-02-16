package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Mysql  MysqlConfig
}

type ServerConfig struct {
	Port        string
	Mode        string
	EnablePprof bool
}

type MysqlConfig struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
}

var Conf Config

func Init(configPath string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		return fmt.Errorf("配置解析失败: %v", err)
	}
	return nil
}
