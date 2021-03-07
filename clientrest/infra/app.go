package infra

import (
	"github.com/kelseyhightower/envconfig"
	"os"
)

type (
	AppConfig struct {
		BasketAddr string `envconfig:"BASKET_SVC_ADDRESS" default:"localhost:8181"`
		Port       string `envconfig:"SVC_PORT" default:":8080"`
	}
)

var (
	App AppConfig
)

func LoadConfig() {
	_ = os.Setenv("LOG_LEVEL", "DEBUG")
	_ = os.Setenv("LOG_FIELDS", "[component=client,service=basket]")
	envconfig.MustProcess("", &App)
}
