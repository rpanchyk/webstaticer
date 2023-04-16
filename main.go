package main

import (
	"log"
	"os"
	"webstaticer/services"
)

func main() {
	log.Println("WebStaticer HTTP Server")

	argsService := services.NewArgsService()
	args, err := argsService.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	configService := services.NewConfigService()
	config, err := configService.Parse(args.ConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	httpService := services.NewHttpService(config)
	err = httpService.Run()
	if err != nil {
		log.Fatal(err)
	}
}
