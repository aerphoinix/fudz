package model

import (
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
)

// type IngredientsList map[*Ingredient]float64
type IngredientsList struct {
	Ingredient Ingredient
	Portion    float64
}

// func (i IngredientsList) MarshallJSON() ([]byte, error) {
// 	return json.Marshal(struct {
// 	}{})
// }

type Recipe struct {
	ID                 uuid.UUID         `json:"id"`
	Name               string            `json:"name"`
	Ingredients        []IngredientsList `json:"ingredients"`
	CaloriesTotal      float64           `json:"calories_total"`
	FatTotal           float64           `json:"fat_total"`
	CarbohydratesTotal float64           `json:"carbohydrates_total"`
	FiberTotal         float64           `json:"fiber_total"`
	ProteinTotal       float64           `json:"protein_total"`
	CostTotal          float64           `json:"cost_total"`
	CostPerFour        float64           `json:"cost_per_four_total"`
	GrossTotal         float64           `json:"gross_total"`
	CreatedAt          *time.Time        `db:"created_at" json:"created_at"`
	UpdatedAt          *time.Time        `db:"updated_at" json:"updated_at"`
}

func NewRecipe(name string, ingredientsList []IngredientsList) *Recipe {
	var caloriesTotal float64
	var fatTotal float64
	var carbohydratesTotal float64
	var fiberTotal float64
	var proteinTotal float64
	var costTotal float64
	var grossCost float64

	for _, ing := range ingredientsList {
		portionPercent := ing.Portion

		caloriesTotal += ing.Ingredient.CaloriesTotal * portionPercent
		fatTotal += ing.Ingredient.FatsTotal * portionPercent
		carbohydratesTotal += ing.Ingredient.CarbohydratesTotal * portionPercent
		fiberTotal += ing.Ingredient.FiberTotal * portionPercent
		proteinTotal += ing.Ingredient.ProteinTotal * portionPercent
		costTotal += ing.Ingredient.CostTotal * portionPercent
		grossCost += ing.Ingredient.CostTotal
	}

	return &Recipe{
		Name:               name,
		Ingredients:        ingredientsList,
		CaloriesTotal:      caloriesTotal,
		FatTotal:           fatTotal,
		CarbohydratesTotal: carbohydratesTotal,
		FiberTotal:         fiberTotal,
		ProteinTotal:       proteinTotal,
		CostTotal:          FixedToTwo(costTotal),
		CostPerFour:        FixedToTwo(costTotal / 4),
		GrossTotal:         FixedToTwo(grossCost),
	}
}

func FixedToTwo(num float64) float64 {
	return math.Round(num*100) / 100
}

func (r *Recipe) Display() {
	fmt.Println("Name: ", r.Name)
	fmt.Println("Ingredients")
	for _, ing := range r.Ingredients {
		fmt.Printf("--> %s | portion: %v\n", ing.Ingredient.Name, ing.Portion)
	}
	fmt.Println("Calories Total:", r.CaloriesTotal)
	fmt.Println("Fat Total:", r.FatTotal)
	fmt.Println("Carbohydrates Total:", r.CarbohydratesTotal)
	fmt.Println("Fiber Total:", r.FiberTotal)
	fmt.Println("Protein Total:", r.ProteinTotal)
	fmt.Println("Cost Total:", r.CostTotal)
	fmt.Println("Cost Per Four:", r.CostPerFour)
	fmt.Println("Gross Total:", r.GrossTotal)
}
