package http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gpbPiazza/garra/infra/envs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newApp(t *testing.T) *fiber.App {
	t.Helper()

	app := NewServer()

	app.Get("/api/v1/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	return app
}

func TestCorsMiddleware(t *testing.T) {
	setDefaultEnvs := func() {
		envsVar := envs.GetEnvs()
		envsVar.AllowOriginsHost = "https://my-frontend.com"
		envsVar.Enviroment = "PRODUCTION"
		envsVar.ApiPort = "8080"
		envs.TestPatchEnvs(envsVar)
	}

	patchEnvAndAllow := func(env, allow string) {
		envsVar := envs.GetEnvs()
		envsVar.AllowOriginsHost = allow
		envsVar.Enviroment = env
		envs.TestPatchEnvs(envsVar)
	}

	setDefaultEnvs()

	t.Run("Development environment should allow all origins", func(t *testing.T) {
		defer setDefaultEnvs()
		patchEnvAndAllow("DEVELOPMENT", "nused-in-dev")

		app := newApp(t)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
		req.Header.Set(fiber.HeaderOrigin, "https://external-site.com")

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, "*", resp.Header.Get(fiber.HeaderAccessControlAllowOrigin))
	})

	t.Run("Production environment should restrict origins", func(t *testing.T) {
		setDefaultEnvs()

		app := newApp(t)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
		req.Header.Set(fiber.HeaderOrigin, "https://my-frontend.com")

		resp, err := app.Test(req)
		require.NoError(t, err)

		require.Equal(t, fiber.StatusOK, resp.StatusCode)
		require.Equal(t, "https://my-frontend.com", resp.Header.Get(fiber.HeaderAccessControlAllowOrigin))
		require.Equal(t, "true", resp.Header.Get(fiber.HeaderAccessControlAllowCredentials))

		// Dissallowed erquest
		req = httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
		req.Header.Set(fiber.HeaderOrigin, "https://malicious-site.com")

		resp, err = app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, "", resp.Header.Get(fiber.HeaderAccessControlAllowOrigin))
	})

	t.Run("Health endpoint should be accessible regardless of CORS", func(t *testing.T) {
		setDefaultEnvs()

		app := newApp(t)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Origin", "https://unauthorized-site.com")

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotEqual(t, "https://unauthorized-site.com", resp.Header.Get("Access-Control-Allow-Origin"))

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, string(body), "I'm Alive")
	})
}

func TestPanicRecoverMiddleware(t *testing.T) {
	app := newApp(t)

	app.Get("panic", func(c *fiber.Ctx) error {
		panic("ai papai!")
	})

	req := httptest.NewRequest(http.MethodGet, "/panic", nil)

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

	req = httptest.NewRequest(http.MethodGet, "/api/v1", nil)
	resp, err = app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)
}
