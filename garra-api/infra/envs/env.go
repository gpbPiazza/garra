package envs

import (
	"log"

	"github.com/caarlos0/env"
)

var globalEnv Env

func init() {
	Init()
}

func Init() {
	var config Env

	if err := env.Parse(&config); err != nil {
		log.Fatalf("error on init envs variables - %s", err)
	}

	globalEnv = config
}

func GetEnvs() Env {
	return globalEnv
}

type Env struct {
	AppName string `env:"NAME" envDefault:"garra-api"`
	ApiPort string `env:"PORT,required"`
}
