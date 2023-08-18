package config

import (
	"github.com/kelseyhightower/envconfig"
)

const APP_PREFIX string = "APP"

type Configuration struct {
	ServerAddr string `envconfig:"SERVER_ADDR" default:":8080"`
}

func New() Configuration {
	var config Configuration

	envconfig.MustProcess(APP_PREFIX, &config)

	return config
}
