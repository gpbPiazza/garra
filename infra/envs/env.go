package envs

import (
	"fmt"
	"log"
	"reflect"

	"github.com/caarlos0/env"
)

type Env struct {
	AppName    string `env:"NAME" envDefault:"garra-api"`
	ApiPort    string `env:"PORT,required"`
	Enviroment string `env:"ENVIROMENT,required"`
}

func GetEnvs() Env {
	valRef := reflect.ValueOf(globalEnv)

	if valRef.IsZero() {
		if err := Init(); err != nil {
			log.Fatal(err)
		}
	}

	return globalEnv
}

var globalEnv Env

func Init() error {
	var config Env

	if err := env.Parse(&config); err != nil {
		return fmt.Errorf("error on init envs variables - %s", err)
	}

	if err := validateEnviroment(config.Enviroment); err != nil {
		return err
	}

	globalEnv = config
	return nil
}

func validateEnviroment(env string) error {
	if env == "PRODUCTION" || env == "DEVELOPMENT" {
		return nil
	}

	return fmt.Errorf("enviroment variable env is not PRODUTCTION OR DEVELOPMENT is - %s", env)
}
