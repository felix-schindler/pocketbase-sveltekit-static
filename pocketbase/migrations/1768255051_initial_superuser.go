package migrations

import (
	"os"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		superusers, err := app.FindCollectionByNameOrId(core.CollectionNameSuperusers)
		if err != nil {
			return err
		}

		email := os.Getenv("ADMIN_MAIL")
		record, _ := app.FindAuthRecordByEmail(core.CollectionNameSuperusers, email)

		if record != nil {
			return nil
		}

		record = core.NewRecord(superusers)
		record.Set("email", email)
		record.Set("password", os.Getenv("ADMIN_PASS"))

		return app.Save(record)
	}, func(app core.App) error {
		record, _ := app.FindAuthRecordByEmail(core.CollectionNameSuperusers, os.Getenv("ADMIN_MAIL"))
		if record == nil {
			return nil
		}

		return app.Delete(record)
	})
}
