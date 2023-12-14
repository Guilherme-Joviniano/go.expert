package main

import "github.com/Guilherme-Joviniano/go.expert/apis/configs"

func main() {
	config, err := configs.LoadConfig(".")
	
	if err != nil {
		panic(err)
	}

	println(config.DBDRiver)

}
