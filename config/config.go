package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 读取yam.yml配置文件
type Options struct {
	Threshold       float64
	ContractAddress string
	UserAddress     string
	PrivateKey      string
	BevmRpc         string
	FhevmHost       string
}

// 读取iam.yml文件，生成options需要的结果
func NewOption(path string) (*Options, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
		return nil, err
	}

	return &Options{
		Threshold:       viper.GetFloat64("threshold"),
		ContractAddress: viper.GetString("contract_address"),
		UserAddress:     viper.GetString("user_address"),
		PrivateKey:      viper.GetString("private_key"),
		BevmRpc:         viper.GetString("bevm_rpc"),
		FhevmHost:       viper.GetString("fhevm_host"),
	}, nil
}

var cfg = &Config{
	Options: &Options{},
}

// Config 自定义配置
type Config struct {
	Options *Options
}

func InitConfig(path string) {
	options, err := NewOption(path)
	if err != nil {
		log.Errorf("解析配置yaml文件失败，错误:[%w]", err)
	}
	cfg.Options = options
}

func GetConfig() *Config {
	return cfg
}
