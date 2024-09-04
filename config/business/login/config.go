package login

import "github.com/zachturing/util/config"

var MapEnvToConfig = map[config.EnvType]config.Param{
	config.DevEnv: {
		APPID:          "login",
		Cluster:        "dev",
		IP:             "http://192.168.200.156:18080",
		Namespace:      "dev",
		IsBackupConfig: true,
		Secret:         "7cb0205ee42b48a88d21192eb5f15ad4",
	},
	config.OnlineEnv: {
		APPID:          "login",
		Cluster:        "dev",
		IP:             "http://192.168.200.156:18080",
		Namespace:      "online",
		IsBackupConfig: true,
		Secret:         "7cb0205ee42b48a88d21192eb5f15ad4",
	},
}
