package main

import (
	"log"
	"os"
	"webstaticer/services"
)

func main() {
	log.Println("WebStaticer HTTP Server")

	argsService := services.NewArgsService(os.Args[1:])
	args, err := argsService.Parse()
	if err != nil {
		log.Fatal(err)
	}

	configService := services.NewConfigService(args.ConfigFile)
	config, err := configService.Parse()
	if err != nil {
		log.Fatal(err)
	}

	httpService := services.NewHttpService(config)
	err = httpService.Run()
	if err != nil {
		log.Fatal(err)
	}
}
