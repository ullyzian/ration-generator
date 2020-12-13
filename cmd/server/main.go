package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/ullyzian/ration-generator/pkg/server"
	"log"
)

var (
	configPath string
)

func init()  {
	flag.StringVar(&configPath, "config-path", "config.toml", "Path to config file")
}

func main() {
	flag.Parse()

	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	serv := server.New(config)
	if err := serv.Start(); err != nil {
		log.Fatal(err)
	}
}

