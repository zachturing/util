package login

import "github.com/zachturing/util/config"

func GetOfficialAccountSubscribeReply() (string, error) {
	cfg, err := config.Get(config.Login)
	if err != nil {
		return "", err
	}
	return cfg.GetStringValue("official_account_subscribe_reply"), nil
}

func GetOfficialAccountLoginReply() (string, error) {
	cfg, err := config.Get(config.Login)
	if err != nil {
		return "", err
	}
	return cfg.GetStringValue("official_account_login_reply"), nil
}
