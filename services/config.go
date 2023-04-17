package services

import (
	"github.com/spf13/viper"
	"log"
	"net"
	"path/filepath"
	"strconv"
	"webstaticer/models"
	"webstaticer/utils"
)

type ConfigService struct {
	filePath string
}

func NewConfigService(filePath string) *ConfigService {
	return &ConfigService{filePath: filePath}
}

func (s *ConfigService) Parse() (*models.Config, error) {
	viper.SetConfigFile(s.filePath)
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

	dir, err := s.normalizeWebRootDir(config.WebRoot.Dir)
	if err != nil {
		return nil, err
	}
	config.WebRoot.Dir = *dir

	port, err := s.normalizeServerPort(config.Server.Port)
	if err != nil {
		return nil, err
	}
	config.Server.Port = *port

	log.Printf("Parsed config:\r\n%s", utils.ToPrettyString(config))

	return &config, nil
}

func (s *ConfigService) normalizeWebRootDir(dir string) (*string, error) {
	path, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	return &path, nil
}

func (s *ConfigService) normalizeServerPort(port string) (*string, error) {
	if port == "0" {
		result, err := s.getFreePort()
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	_, err := strconv.ParseInt(port, 10, 64)
	if err == nil {
		return nil, err
	}
	return &port, nil
}

func (s *ConfigService) getFreePort() (*string, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return nil, err
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	port := listener.Addr().(*net.TCPAddr).Port
	portAsString := strconv.Itoa(port)

	return &portAsString, nil
}
