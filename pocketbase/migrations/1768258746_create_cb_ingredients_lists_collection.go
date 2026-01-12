package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection := core.NewBaseCollection("cb_ingredients_lists")

		user_authenticated := "@request.auth.id != null"
		collection.ListRule = &user_authenticated
		collection.ViewRule = &user_authenticated
		collection.CreateRule = &user_authenticated
		user_is_author := "@request.auth.id = author"
		collection.UpdateRule = &user_is_author
		collection.DeleteRule = &user_is_author

		recipe, _ := app.FindCollectionByNameOrId("recipe")
		author, _ := app.FindCollectionByNameOrId("users")

		collection.Fields.Add(
			&core.TextField{
				Name:     "title",
				Required: true,
			},
			&core.RelationField{
				Name:          "recipe",
				CollectionId:  recipe.Id,
				CascadeDelete: true,
				Required:      true,
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
		collection, err := app.FindCollectionByNameOrId("cb_ingredients_lists")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
