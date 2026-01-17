package main

import (
	"log"

	"github.com/hisbaan/envman/cmd"
	"github.com/hisbaan/envman/config"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}
	cmd.Execute()
}
