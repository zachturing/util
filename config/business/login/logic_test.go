package login

import (
	"testing"

	"github.com/zachturing/util/config"
)

func TestGetOfficialAccountSubscribeReply(t *testing.T) {
	err := config.Register(config.Login, MapEnvToConfig, config.DevEnv)
	if err != nil {
		t.Errorf("register failed, err:%v", err)
		return
	}

	reply, err := GetOfficialAccountSubscribeReply()
	if err != nil {
		t.Errorf("get config failed, err:%v", err)
		return
	}
	t.Logf("reply :%v", reply)
}

func TestGetOfficialAccountLoginReply(t *testing.T) {
	err := config.Register(config.Login, MapEnvToConfig, config.DevEnv)
	if err != nil {
		t.Errorf("register failed, err:%v", err)
		return
	}

	reply, err := GetOfficialAccountLoginReply()
	if err != nil {
		t.Errorf("get config failed, err:%v", err)
		return
	}
	t.Logf("reply :%v", reply)
}
