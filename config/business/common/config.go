package common

import "github.com/zachturing/util/config"

var MapEnvToConfig = map[config.EnvType]config.Param{
	config.DevEnv: {
		APPID:          "common",
		Cluster:        "DEV",
		IP:             "http://124.222.11.142:8080",
		Namespace:      "mix-paper.dev",
		IsBackupConfig: true,
		// Secret:         "c05782a4367d4198b660b31c766862f6",
	},
	config.OnlineEnv: {
		APPID:          "common",
		Cluster:        "DEV",
		IP:             "http://124.222.11.142:18080",
		Namespace:      "mix-paper.prod",
		IsBackupConfig: true,
		// Secret:         "c05782a4367d4198b660b31c766862f6",
	},
}
