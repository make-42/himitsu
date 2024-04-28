package config

import (
	"himitsu/totp"
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
	"gopkg.in/yaml.v3"
)

const Version = "v0.0.1"

var DefaultConfig = []totp.TOTP{
	{
		Label:   "Label",
		Secret:  "Secret",
		Account: "Account",
	},
}
var Config []totp.TOTP

func Init() {
	configPath := configdir.LocalConfig("ontake", "himitsu")
	err := configdir.MakePath(configPath) // Ensure it exists.
	if err != nil {
		panic(err)
	}

	configFile := filepath.Join(configPath, "totp.yml")

	// Does the file not exist?
	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		// Create the new config file.
		fh, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer fh.Close()

		encoder := yaml.NewEncoder(fh)
		encoder.Encode(&DefaultConfig)
		Config = DefaultConfig
	} else {
		// Load the existing file.
		fh, err := os.Open(configFile)
		if err != nil {
			panic(err)
		}
		defer fh.Close()

		decoder := yaml.NewDecoder(fh)
		decoder.Decode(&Config)
	}
}
