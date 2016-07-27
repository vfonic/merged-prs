package config

import (
	"log"
	"os/user"
)

const ConfigFile = ".merged-prs"

type Config struct {
	User string
	Home string
}

func NewDefault() *Config {
	usr, err := user.Current()
	if err != nil {
		// TODO: logger
		log.Fatal("Cannot get user")
	}

	return &Config{
		User: usr.Username,
		Home: usr.HomeDir,
	}
}

func NewMock() *Config {
	return &Config{
		User: "test",
		Home: ".",
	}
}
