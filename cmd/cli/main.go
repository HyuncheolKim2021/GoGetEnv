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
	fmt.Printf("Load Database Env with `go:embed`: %#v\n", cfg2.Database)

	cfg3, err := config.LoadConfigJsonWithEmbed("local")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Load Database Env with `go:embed`: %#v\n", cfg3.Database)

	cfg4, err := config.LoadConfigYamlWithEmbed("local")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Load Database Env with `go:embed`: %#v\n", cfg4.Database)

}
