package services

import (
	"log"
	"net"
	"net/http"
	"path/filepath"
	"strconv"
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

	port, err := s.getPort()
	if err != nil {
		return err
	}
	portAsString := strconv.Itoa(port)

	path, err := filepath.Abs(s.config.WebRoot.Dir)
	if err != nil {
		return err
	}
	log.Printf("Running http://%s:%s for \"%s\"", host, portAsString, path)
	log.Println("Press Ctrl+C to exit...")

	http.HandleFunc("/", s.serveFiles)
	err = http.ListenAndServe(host+":"+portAsString, nil)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (s *HttpService) getPort() (int, error) {
	port := s.config.Server.Port

	if port == "0" {
		free, err := getFreePort()
		if err != nil {
			return -1, err
		}
		return free, nil
	}

	parsed, err := strconv.ParseInt(port, 10, 64)
	if err == nil {
		return -1, err
	}
	return int(parsed), nil
}

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return -1, err
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return -1, err
	}

	defer func() {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return listener.Addr().(*net.TCPAddr).Port, nil
}

func (s *HttpService) serveFiles(writer http.ResponseWriter, request *http.Request) {
	urlPath := request.URL.Path
	if urlPath == "/" {
		urlPath = s.config.WebRoot.File
	}
	http.ServeFile(writer, request, filepath.Join(s.config.WebRoot.Dir, urlPath))
}
