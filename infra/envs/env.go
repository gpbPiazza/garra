package envs

import (
	"fmt"
	"log"
	"reflect"

	"github.com/caarlos0/env"
)

type Env struct {
	AppName          string `env:"NAME" envDefault:"garra-api"`
	ApiPort          string `env:"PORT,required"`
	Enviroment       string `env:"ENVIROMENT,required"`
	AllowOriginsHost string `env:"ALLOW_ORIGINS_HOST,required"`
}

func GetEnvs() Env {
	valRef := reflect.ValueOf(globalEnv)

	if valRef.IsZero() {
		if err := initEnvs(); err != nil {
			log.Fatal(err)
		}
	}

	return globalEnv
}

func TestPatchEnvs(env Env) Env {
	globalEnv = env
	return globalEnv
}

func IsProduction() bool {
	return GetEnvs().Enviroment == "PRODUCTION"
}

var globalEnv Env

func initEnvs() error {
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
