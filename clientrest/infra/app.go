package infra

import "github.com/kelseyhightower/envconfig"

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
	envconfig.MustProcess("", &App)
}
