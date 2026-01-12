package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection := core.NewBaseCollection("cb_recipes")

		user_authenticated := "@request.auth.id != null"
		collection.ListRule = &user_authenticated
		collection.ViewRule = &user_authenticated
		collection.CreateRule = &user_authenticated
		user_is_author := "@request.auth.id = author"
		collection.UpdateRule = &user_is_author
		collection.DeleteRule = &user_is_author

		cuisine, _ := app.FindCollectionByNameOrId("cb_cuisines")
		tags, _ := app.FindCollectionByNameOrId("cb_tags")
		author, _ := app.FindCollectionByNameOrId("users")

		collection.Fields.Add(
			&core.TextField{
				Name:     "title",
				Required: true,
			},
			&core.TextField{
				Name: "description",
			},
			&core.NumberField{
				Name: "prepTimeInMinutes",
			},
			&core.NumberField{
				Name: "cookTimeInMinutes",
			},
			&core.NumberField{
				Name:     "servings",
				Required: true,
			},
			&core.SelectField{
				Name:     "difficulty",
				Values:   []string{"easy", "medium", "hard"},
				Required: true,
			},
			&core.FileField{
				Name: "image",
			},
			&core.URLField{
				Name: "altImageUrl",
			},
			&core.RelationField{
				Name:          "cuisine",
				CollectionId:  cuisine.Id,
				CascadeDelete: false,
				Required:      true,
			},
			&core.RelationField{
				Name:          "tags",
				CollectionId:  tags.Id,
				MaxSelect:     999,
				CascadeDelete: false,
			},
			&core.RelationField{
				Name:          "author",
				CollectionId:  author.Id,
				CascadeDelete: false,
				Required:      true,
			},
		)

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("cb_recipes")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
