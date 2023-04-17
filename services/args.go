package services

import (
	"log"
	"os"
	"path/filepath"
	"webstaticer/models"
	"webstaticer/utils"
)

const defaultConfigFile = "webstaticer.yaml"

type ArgsService struct {
	inputArguments []string
}

func NewArgsService(inputArguments []string) *ArgsService {
	return &ArgsService{inputArguments: inputArguments}
}

func (s *ArgsService) Parse() (*models.Args, error) {
	configFile, err := s.getConfigFile()
	if err != nil {
		return nil, err
	}

	args := &models.Args{ConfigFile: *configFile}
	log.Printf("Parsed arguments:\r\n%s", utils.ToPrettyString(args))

	return args, nil
}

func (s *ArgsService) getConfigFile() (*string, error) {
	configFile := defaultConfigFile
	if len(s.inputArguments) > 0 {
		configFile = s.inputArguments[0]
	}

	configFilePath, err := filepath.Abs(configFile)
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(configFilePath)
	if err != nil {
		return nil, err
	}

	return &configFilePath, nil
}
