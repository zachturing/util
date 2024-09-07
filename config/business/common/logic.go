package common

import (
	oaconfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/zachturing/util/config"
	"github.com/zachturing/util/database"
	"github.com/zachturing/util/sms"
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

func GetSMSConfig() (*sms.Param, error) {
	cfg, err := config.Get(config.Common)
	if err != nil {
		return nil, err
	}
	var smsConfig sms.Param
	err = cfg.GetWithUnmarshal("sms", &smsConfig, &config.JSONUnmarshaler{})
	if err != nil {
		return nil, err
	}
	return &smsConfig, nil
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
