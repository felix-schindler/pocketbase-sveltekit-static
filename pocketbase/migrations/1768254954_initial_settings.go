package migrations

import (
	"os"
	"strconv"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	// Docs https://pocketbase.io/docs/go-migrations/#initialize-default-application-settings
	m.Register(func(app core.App) error {
		settings := app.Settings()

		// for all available settings fields you could check
		// https://github.com/pocketbase/pocketbase/blob/develop/core/settings_model.go#L121-L130
		settings.Meta.AppName = "Hub"
		settings.Meta.AppURL = "https://pb-sk.schindlerfelix.de/"
		settings.Logs.MaxDays = 3
		settings.Logs.LogAuthId = true
		settings.Logs.LogIP = true

		settings.TrustedProxy.Headers = []string{"CF-Connecting-IP"}
		settings.TrustedProxy.UseLeftmostIP = false
		settings.RateLimits.Enabled = true

		settings.Meta.SenderName = os.Getenv("SMTP_SENDER_NAME")
		settings.Meta.SenderAddress = os.Getenv("SMTP_SENDER_ADDRESS")
		settings.SMTP.Host = os.Getenv("SMTP_HOST")
		settings.SMTP.Port, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
		settings.SMTP.Username = os.Getenv("SMTP_USERNAME")
		settings.SMTP.Password = os.Getenv("SMTP_PASSWORD")
		settings.SMTP.Enabled = true

		return app.Save(settings)
	}, nil)
}
