package config

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

//go:embed default.yaml
var defaultYaml []byte

type ServerConfig struct {
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	LogLevel string `mapstructure:"log_level"`
}

type DbConfig struct {
	Type     string `mapstructure:"type"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type AppConfig struct {
	Server ServerConfig `mapstructure:"server"`
	Db     DbConfig     `mapstructure:"db"`
}

var Cfg AppConfig

func InitConfig() {
	// 1. 处理命令行参数
	var configFile string
	pflag.StringVar(&configFile, "config", "", "Path to custom config file")
	pflag.IntVar(&Cfg.Server.Port, "server.port", 8888, "Server port")
	pflag.Parse()

	v := viper.New()

	// 2. 加载内嵌的默认配置
	v.SetConfigType("yaml")
	if err := v.ReadConfig(strings.NewReader(string(defaultYaml))); err != nil {
		fmt.Println("加载默认配置失败:", err)
		os.Exit(1)
	}

	// 3. 如果有外部配置文件，就加载并合并
	if configFile != "" {
		v.SetConfigFile(configFile)
		if err := v.MergeInConfig(); err != nil {
			fmt.Println("加载外部配置失败:", err)
			os.Exit(1)
		}
	}

	// 4. 环境变量覆盖（支持 HERTZ_SERVER_PORT 这类变量）
	v.SetEnvPrefix("HERTZ")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// 5. 合并命令行参数
	_ = v.BindPFlags(pflag.CommandLine)

	// 6. 映射到结构体
	if err := v.Unmarshal(&Cfg); err != nil {
		fmt.Println("解析配置失败:", err)
		os.Exit(1)
	}
}
