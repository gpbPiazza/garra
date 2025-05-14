package http

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCorsMiddleware(t *testing.T) {
	setDefaultEnvs := func() {
		os.Setenv("PORT", "8080")
		os.Setenv("ENVIROMENT", "PRODUCTION")
		os.Setenv("ALLOW_ORIGINS_HOST", "https://my-frontend.com")
	}

	setDefaultEnvs()

	newApp := func() *fiber.App {
		app := NewServer()

		app.Get("/api/v1/", func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		return app
	}

	t.Run("Development environment should allow all origins", func(t *testing.T) {
		os.Setenv("ENVIROMENT", "DEVELOPMENT")
		os.Setenv("ALLOW_ORIGINS_HOST", "unused-in-dev")

		app := newApp()

		req := httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
		req.Header.Set(fiber.HeaderOrigin, "https://external-site.com")

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, "*", resp.Header.Get(fiber.HeaderAccessControlAllowOrigin))
	})

	t.Run("Production environment should restrict origins", func(t *testing.T) {
		setDefaultEnvs()

		app := newApp()

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

	// t.Run("Health endpoint should be accessible regardless of CORS", func(t *testing.T) {
	// 	setDefaultEnvs()

	// 	app := newApp()

	// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// 	req.Header.Set("Origin", "https://unauthorized-site.com")

	// 	resp, err := app.Test(req)
	// 	assert.NoError(t, err)

	// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
	// 	assert.NotEqual(t, "https://unauthorized-site.com", resp.Header.Get("Access-Control-Allow-Origin"))

	// 	body, err := io.ReadAll(resp.Body)
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, string(body), "I'm Alive")
	// })
}
