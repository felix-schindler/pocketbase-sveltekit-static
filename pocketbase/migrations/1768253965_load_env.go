package migrations

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		if os.Getenv("APP_ENV") != "production" {
			godotenv.Load()
		}

		return nil
	}, nil)
}
