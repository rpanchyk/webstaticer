package services

import (
	"log"
	"net/http"
	"path/filepath"
	"webstaticer/models"
)

type HttpService struct {
	config *models.Config
}

func NewHttpService(config *models.Config) *HttpService {
	return &HttpService{config: config}
}

func (s *HttpService) Run() error {
	host := s.config.Server.Host
	port := s.config.Server.Port

	log.Printf("Running http://%s:%s for \"%s\"", host, port, s.config.WebRoot.Dir)
	log.Println("Press Ctrl+C to exit...")

	http.HandleFunc("/", s.serveFiles)
	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *HttpService) serveFiles(writer http.ResponseWriter, request *http.Request) {
	urlPath := request.URL.Path
	if urlPath == "/" {
		urlPath = s.config.WebRoot.File
	}
	path := filepath.Join(s.config.WebRoot.Dir, urlPath)

	http.ServeFile(writer, request, path)
}
