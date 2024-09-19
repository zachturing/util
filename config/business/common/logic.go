package common

import (
	oaconfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/zachturing/util/config"
	"github.com/zachturing/util/database"
)

func getConfig(key string) (*database.DBConfig, error) {
	cfg, err := config.Get(config.Common)
	if err != nil {
		return nil, err
	}
	var dbConfig database.DBConfig
	err = cfg.GetWithUnmarshal(key, &dbConfig, &config.JSONUnmarshaler{})
	if err != nil {
		return nil, err
	}
	return &dbConfig, nil
}

func GetMysqlConfig() (*database.DBConfig, error) {
	return getConfig("mysql")
}

func GetRedisConfig() (*database.DBConfig, error) {
	return getConfig("redis")
}

// Param 短信参数
type Param struct {
	Signature  string `json:"signature"`
	TemplateID string `json:"template_id"`
	TTL        string `json:"ttl"`
}

func GetSMSConfig() (*Param, error) {
	cfg, err := config.Get(config.Common)
	if err != nil {
		return nil, err
	}
	var param Param
	err = cfg.GetWithUnmarshal("sms", &param, &config.JSONUnmarshaler{})
	if err != nil {
		return nil, err
	}
	return &param, nil
}

func GetOfficialAccountConfig() (*oaconfig.Config, error) {
	cfg, err := config.Get(config.Common)
	if err != nil {
		return nil, err
	}

	oaCfg := &oaconfig.Config{}
	err = cfg.GetWithUnmarshal("wechat_official_account", oaCfg, &config.JSONUnmarshaler{})
	if err != nil {
		return nil, err
	}
	return oaCfg, nil
}
