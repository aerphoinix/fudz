package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aerphoinix/fudz/internal/database"
	"github.com/aerphoinix/fudz/internal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// initialize sqlx pool
	pgPool, err := database.NewPGPool()
	if err != nil {
		log.Fatal(err)
	}

	// new echo instance
	e := echo.New()

	// beans := model.NewIngredient("beans", "stater brothers", 3.5, []string{"can"}, 454, 110, 2.5, 16, 7, 6, 1.39)
	// tortillas := model.NewIngredient("tortills", "stater brothers", 3.5, []string{"bag"}, 1000, 66, 2, 10, 0, 1, 2.48)
	// cheese := model.NewIngredient("cheese", "stater brothers", 3.5, []string{"bag"}, 454, 140, 9, 1, 0, 8, 5.50)

	torrilas := model.Ingredient{}
	err = pgxscan.Get(context.Background(), pgPool, &torrilas, "SELECT * FROM ingredient WHERE id = $1", "ce2fcb57-7700-437a-8043-c633be749cf2")
	fmt.Println(err)

	// tacos := model.NewRecipe("Tacos", []model.IngredientsList{{Ingredient: beans, Portion: 1}, {Ingredient: *tortillas, Portion: 0.4}, {Ingredient: *cheese, Portion: 0.5}})

	// tacos.Display()

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })

	// start echo
	// e.Logger.Fatal(e.Start(":1323"))

	// InsertIngredient(beans, pgPool)
	// InsertRecipe(tacos, pgPool)

	// fmt.Println(err)
	// tacos.Display()

	e.GET("/api/recipes", func(c echo.Context) error {
		recipes := []model.Recipe{}
		err := pgxscan.Select(context.Background(), pgPool, &recipes, "SELECT * FROM recipe")
		if err != nil {
			return c.JSON(404, err)
		}

		return c.JSON(200, recipes)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func InsertIngredient(ingredient *model.Ingredient, pgPool *database.PGPool) error {
	var id uuid.UUID
	err := pgPool.QueryRow(context.Background(), "INSERT INTO ingredient (name, brand, servings, packaging, calories_per_serving, calories_total, fats_per_serving, fats_total, carbohydrates_per_serving, carbohydrates_total, fiber_per_serving, fiber_total, protein_per_serving, protein_total, cost_per_serving, cost_total) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id", &ingredient.Name, &ingredient.Brand, &ingredient.Servings, &ingredient.Packaging, &ingredient.CaloriesPerServing, &ingredient.CaloriesTotal, &ingredient.FatsPerServing, &ingredient.FatsTotal, &ingredient.CarbohydratesPerServing, &ingredient.CarbohydratesTotal, &ingredient.FiberPerServing, &ingredient.FiberTotal, &ingredient.ProteinPerServing, &ingredient.ProteinTotal, &ingredient.CostPerServing, &ingredient.CostTotal).Scan(&id)

	return err
}

func InsertRecipe(recipe *model.Recipe, pgPool *database.PGPool) error {
	var id uuid.UUID

	ingredientJson, err := json.Marshal(recipe.Ingredients)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = pgPool.QueryRow(context.Background(), "INSERT INTO recipe (name, ingredients, calories_total, fat_total, carbohydrates_total, fiber_total, protein_total, cost_total, cost_per_four, gross_total) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id", &recipe.Name, &ingredientJson, &recipe.CaloriesTotal, &recipe.FatTotal, &recipe.CarbohydratesTotal, &recipe.FiberTotal, &recipe.ProteinTotal, &recipe.CostTotal, &recipe.CostPerFour, &recipe.GrossTotal).Scan(&id)

	fmt.Println(id)

	return err
}
