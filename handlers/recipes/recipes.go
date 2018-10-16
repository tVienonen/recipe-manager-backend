package recipes

import (
	"time"

	"github.com/kataras/iris"
	"github.com/tVienonen/recipe-manager-backend/resources/database"
)

type Recipe struct {
	ID          *uint64    `json:"id,omitempty"`
	Name        *string    `json:"name" validate:"required"`
	Description *string    `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	ModifiedAt  *time.Time `json:"modified_at,omitempty"`
	UserID      *time.Time `json:"user_id,omitempty"`
}

type RecipePicture struct {
	ID          *uint64    `json:"id,omitempty"`
	URI         *string    `json:"uri,omitempty"`
	Description *string    `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	ModifiedAt  *time.Time `json:"modified_at,omitempty"`
	RecipeID    *uint64    `json:"recipe_id,omitempty"`
}

func GetRecipesHandler(ctx iris.Context) {
	db := database.GetDB(ctx)
	recipes := make([]Recipe, 0)
	q := db.NewQuery("SELECT * FROM recipes;")
	if err := q.All(&recipes); err != nil {
		panic(err)
	}
	ctx.JSON(recipes)
}
