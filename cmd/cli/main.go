package main

import (
	"fmt"
	"vipersample/config"
)

func main() {
	cfg, err := config.LoadConfigWithViper("local")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Load Database Env with Viper: %#v\n", cfg.Database)

	cfg2, err := config.LoadConfigWithEmbed("local")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Load Database Env with `go:embed`: %#v", cfg2.Database)
}
