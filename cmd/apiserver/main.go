package main

import (
	"flag"
	"log"

	"github.com/AveCaesar17/basic-server-go.git/internal/app/apiserver"
	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "Путь до конфиг файла ")
}
func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	server := apiserver.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
