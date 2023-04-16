package services

import (
	"github.com/spf13/viper"
	"log"
	"webstaticer/models"
	"webstaticer/utils"
)

type ConfigService struct {
}

func NewConfigService() *ConfigService {
	return &ConfigService{}
}

func (s *ConfigService) Parse(filePath string) (*models.Config, error) {
	viper.SetConfigFile(filePath)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config models.Config

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	log.Printf("Parsed \"%s\" config:\r\n%s", filePath, utils.ToPrettyString(config))

	return &config, nil
}
