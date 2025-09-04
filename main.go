package main

import (
	"github.com/dulchik/blog_aggregator/internal/config"

	"fmt"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := cfg.SetUser("dulchik"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(cfg)
}

