package main

import (
	"fmt"
	"log"

	"github.com/rwndy/go-blog/internal/config"
)

func main() {
	// read config
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	// set user
	err = cfg.SetUser("rwndy")
	if err != nil {
		log.Fatal(err)
	}

	// read config again
	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", cfg)
}
