package relay

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Config is the configuration for the relay.
type Config struct {
	// Filter defines a function to skip middleware.
	// Optional. Default: nil
	Filter func(*fiber.Ctx) bool

	// HandshakeTimeout specifies the duration for the handshake to complete.
	HandshakeTimeout time.Duration
}

// New creates a new relay middleware.
func New(config Config, addr string) fiber.Handler {
	// Set default config
	cfg := configDefault(config)

	return func(c *fiber.Ctx) error {
		if cfg.Filter != nil && cfg.Filter(c) {
			return c.Next()
		}

		return nil
	}
}

func configDefault(config ...Config) Config {
	// Init config
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}

	return cfg
}
