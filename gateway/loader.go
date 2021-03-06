package gateway

import (
	"github.com/go-playground/validator"
	"github.com/ikeikeikeike/cheapcdn/config"

	"github.com/labstack/echo"
	md "github.com/labstack/echo/middleware"
)

var (
	cfg   *config.Config
	valid = validator.New()
)

// Load configuration instead of context.
func Load(c *config.Config) {
	cfg = c
}

// Routes set handler into mux
func Routes(e *echo.Echo) {
	g := e.Group("/gateway")
	g.Use(md.BasicAuth(basicAuth))

	g.GET("/", gateway)
}

func init() {
	valid = validator.New()
}
