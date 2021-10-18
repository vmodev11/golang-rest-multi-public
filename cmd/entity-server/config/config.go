package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var cf *Configuration

type Configuration struct {
	EnvironmentPrefix string
	DbConnection      string
	AllowOrigins      []string
}

func GetConfig() *Configuration {
	return cf
}

func InitFromFile(filePathStr string, basePath string) {

	env := os.Getenv("GO_ENV")

	viper.SetConfigFile(filePathStr)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Config file not found: %v", err)
	} else {

		environmentPrefix := viper.GetString(env + ".environment_prefix")
		cf = &Configuration{
			EnvironmentPrefix: environmentPrefix,
			DbConnection:      viper.GetString(env + ".db_connection"),
			AllowOrigins:      viper.GetStringSlice(env + ".allow_origins"),
		}
		log.Println(viper.ConfigFileUsed())
		log.Printf("Config %+v", *cf)
	}
}
