package conf

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

func LoadConfigFile() error {
	//Read the configuration file from environment variable and provide application wide access
	filePath := os.Getenv("TASKLOGGER_CONF_FILE")

	log.Println(filePath)

	viper.AddConfigPath(filePath)
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		log.Println("Could not find TASKLOGGER_CONF_FILE enviroment variable, which should point to conf directory path")
		return err
	}
	return nil

}
