package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config Struct para configurações do servidor e do banco de dados
type Config struct {
	Server   string
	Database string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
