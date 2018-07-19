package main

import (
	"Neo-Rank/config"
	"fmt"
)

const (
	configPath = "./config"
	env        = "development"
)

func main() {
	fmt.Println("main.go")
	cfg, err := config.Load(configPath, env)
	if err != nil {

	}
}
