package settings

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Environments : Maps the environment files
var Environments = map[string]string{
	"dev": "settings/dev.json",
}

// default environment
var env = "dev"

// Init : Init function that is called to initialize the settings file
func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting dev environment due to lack of GO_ENV value")
		env = "dev"
	}
	LoadSettingsByEnv(env)
}

// LoadSettingsByEnv : loads the environments by setting
func LoadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(Environments[env])
	if err != nil {
		fmt.Println("Error while reading config file", err)
	}
	fmt.Print(content)
}
