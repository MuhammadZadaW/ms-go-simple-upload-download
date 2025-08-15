package configs

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type ConfigStruct struct {
	Server struct {
		Port int    `mapstructure:"PORT"`
		Name string `mapstructure:"NAME"`
	} `mapstructure:"SERVER"`
	Upload struct {
		DestinationPath string `mapstructure:"DESTINATION_PATH"`
	} `mapstructure:"UPLOAD"`
	Download struct {
		SourcePath string `mapstructure:"SOURCE_PATH"`
	} `mapstructure:"DOWNLOAD"`
}

func LoadConfig() (*ConfigStruct, error) {

	var config ConfigStruct

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Failed to read configuration file")
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Println("Failed to unmarshal configuration")
		return nil, err
	}

	_, err = os.Stat(config.Upload.DestinationPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(config.Upload.DestinationPath, os.ModePerm)
		if err != nil {
			log.Println("Failed to create upload directory")
			return nil, err
		}
	}

	_, err = os.Stat(config.Download.SourcePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(config.Download.SourcePath, os.ModePerm)
		if err != nil {
			log.Println("Failed to create download directory")
			return nil, err
		}
	}

	return &config, nil
}
