package main

import (
	"log"

	"github.com/nuea/go-template/cmd/http/di"
)

func main() {
	ctn, stop, err := di.InitContainer()
	if err != nil {
		log.Panicf("Unable to start service. Error: %s", err)
	}
	defer stop()
	ctn.Run()
}
