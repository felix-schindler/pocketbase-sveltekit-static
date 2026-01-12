package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection := core.NewBaseCollection("cb_ingredients")

		user_authenticated := "@request.auth.id != null"
		collection.ListRule = &user_authenticated
		collection.ViewRule = &user_authenticated
		collection.CreateRule = &user_authenticated
		user_is_author := "@request.auth.id = author"
		collection.UpdateRule = &user_is_author
		collection.DeleteRule = &user_is_author

		ingredientsList, _ := app.FindCollectionByNameOrId("ingredientsList")
		author, _ := app.FindCollectionByNameOrId("users")

		collection.Fields.Add(
			&core.TextField{
				Name:     "name",
				Required: true,
			},
			&core.NumberField{
				Name: "quantity",
			},
			&core.SelectField{
				Name:   "unit",
				Values: []string{"Tl", "Tsp", "El", "ml", "l", "mg", "g", "kg", "pcs", "Can"},
			},
			&core.RelationField{
				Name:          "ingredientsList",
				CollectionId:  ingredientsList.Id,
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
		collection, err := app.FindCollectionByNameOrId("cb_ingredients")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
