package envs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	os.Setenv("PORT", "8080")
	os.Setenv("ENVIROMENT", "DEVELOPMENT")
	os.Setenv("NAME", "garra-dev")
	os.Setenv("ALLOW_ORIGINS_HOST", "LOCALHOST")

	originalPort := os.Getenv("PORT")
	originalEnv := os.Getenv("ENVIROMENT")
	originalName := os.Getenv("NAME")
	originalAllowOriginsHost := os.Getenv("ALLOW_ORIGINS_HOST")

	defer func() {
		os.Setenv("PORT", originalPort)
		os.Setenv("ENVIROMENT", originalEnv)
		os.Setenv("NAME", originalName)
		os.Setenv("ALLOW_ORIGINS_HOST", originalAllowOriginsHost)
		initEnvs()
	}()

	t.Run("should initialize with PRODUCTION environment", func(t *testing.T) {
		os.Setenv("PORT", "8080")
		os.Setenv("ENVIROMENT", "PRODUCTION")

		globalEnv = Env{}

		err := initEnvs()

		assert.NoError(t, err)
		assert.Equal(t, "PRODUCTION", globalEnv.Enviroment)
		assert.Equal(t, "8080", globalEnv.ApiPort)
		assert.Equal(t, "garra-dev", globalEnv.AppName)
	})

	t.Run("should initialize with DEVELOPMENT environment", func(t *testing.T) {
		os.Setenv("PORT", "3000")
		os.Setenv("ENVIROMENT", "DEVELOPMENT")
		os.Setenv("NAME", "garra-ZAP")

		globalEnv = Env{}

		err := initEnvs()

		assert.NoError(t, err)
		assert.Equal(t, "DEVELOPMENT", globalEnv.Enviroment)
		assert.Equal(t, "3000", globalEnv.ApiPort)
		assert.Equal(t, "garra-ZAP", globalEnv.AppName)
	})

	t.Run("should fail with invalid environment", func(t *testing.T) {
		os.Setenv("PORT", "8080")
		os.Setenv("ENVIROMENT", "STAGING")

		globalEnv = Env{}

		err := initEnvs()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not PRODUTCTION OR DEVELOPMENT")
	})

	t.Run("should fail when required env vars are missing", func(t *testing.T) {
		os.Unsetenv("PORT")
		os.Unsetenv("ENVIROMENT")

		globalEnv = Env{}

		err := initEnvs()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error on init envs variables")
	})
}
