package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

type AppConfig struct {
	Server ServerConfig `mapstructure:"server"`
	// 可扩展更多模块配置
}

var Cfg AppConfig

func InitConfig() {
	// 1. 命令行参数
	var configFile string
	pflag.StringVar(&configFile, "config", "config/config.yaml", "Path to custom config file")
	pflag.IntVar(&Cfg.Server.Port, "server.port", 8888, "Server port")
	pflag.Parse()

	// 2. 默认配置文件
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	_ = viper.MergeInConfig()

	// 3. 用户自定义配置文件（覆盖默认）
	if configFile != "" {
		viper.SetConfigFile(configFile)
		_ = viper.MergeInConfig()
	}

	// 4. 绑定环境变量（支持下划线转小写点）
	viper.SetEnvPrefix("HERTZ")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// 5. pflag 合并
	_ = viper.BindPFlags(pflag.CommandLine)

	// 6. 映射到结构体
	if err := viper.Unmarshal(&Cfg); err != nil {
		fmt.Println("读取配置失败:", err)
		os.Exit(1)
	}
}
