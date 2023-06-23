package libutils

import (
	"fmt"

	"github.com/spf13/viper"
)

// init configuration for project by reading file config/app.yaml
func InitConfig(config_dir string) (err error) {
	viper.SetConfigFile(config_dir)
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// get value associated to a key in file config
// key can be nested, use a dot to seperate level
func Get(key string) string {
	return viper.GetString(key)
}
