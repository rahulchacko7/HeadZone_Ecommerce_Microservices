package main

import (
	"log"

	"github.com/rahulchacko7/GO-GRPC-ORDER-SVC/pkg/config"
	"github.com/rahulchacko7/GO-GRPC-ORDER-SVC/pkg/di"
)

func main() {

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)

	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}

}
