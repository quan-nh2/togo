package main

import (
	"fmt"
	"togo/config"
)

func main() {
	cfg, err := config.ParseConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg)
}
