package config

import "himitsu/totp"

const Version = "v0.0.1"

var Config = []totp.TOTP{
	{
		Label:   "Test",
		Secret:  "Test",
		Account: "Test",
	},
}
