package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection := core.NewBaseCollection("cb_cuisines")

		user_authenticated := "@request.auth.id != null"
		collection.ListRule = &user_authenticated
		collection.ViewRule = &user_authenticated

		collection.Fields.Add(
			&core.TextField{
				Name:     "title",
				Required: true,
			},
			&core.TextField{
				Name:     "flag",
				Required: true,
			},
			&core.FileField{
				Name: "image",
			},
		)

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("cb_cuisines")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
