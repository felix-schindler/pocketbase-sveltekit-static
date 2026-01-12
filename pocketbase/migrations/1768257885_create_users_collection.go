package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection := core.NewAuthCollection("users")

		user_himself := "id = @request.auth.id"
		collection.ListRule = &user_himself
		collection.ViewRule = &user_himself
		collection.UpdateRule = &user_himself
		collection.DeleteRule = &user_himself

		features := []string{"cookbook", "experiments", "social"}

		collection.Fields.Add(
			&core.TextField{
				Name:     "name",
				Required: true,
				Max:      255,
			},
			&core.FileField{
				Name:      "avatar",
				MimeTypes: []string{"image/jpeg", "image/png", "image/svg+xml", "image/gif", "image/webp"},
			},
			&core.SelectField{
				Name:   "role",
				Values: []string{"admin", "user"},
			},
			&core.SelectField{
				Name:      "features",
				Values:    features,
				MaxSelect: len(features),
			},
		)

		collection.PasswordAuth.Enabled = true

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
