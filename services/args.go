package services

import (
	"log"
	"os"
	"webstaticer/models"
	"webstaticer/utils"
)

const defaultConfigFile = "webstaticer.yaml"

type ArgsService struct {
}

func NewArgsService() *ArgsService {
	return &ArgsService{}
}

func (s *ArgsService) Parse(args []string) (*models.Args, error) {
	configFile := defaultConfigFile
	if len(args) > 0 {
		configFile = args[0]
	}

	_, err := os.Stat(configFile)
	if err != nil {
		return nil, err
	}

	result := &models.Args{ConfigFile: configFile}
	log.Printf("Parsed arguments:\r\n%s", utils.ToPrettyString(result))

	return result, nil
}
