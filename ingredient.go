package main

import (
	"math"
	"time"

	"github.com/google/uuid"
)

type Ingredient struct {
	ID                      uuid.UUID `json:"id"`
	Name                    string    `json:"name"`
	Brand                   string    `json:"brand"`
	Servings                float64   `json:"servings"`
	Packaging               []string  `json:"packaging"`
	WeightTotalOunces       float64   `json:"weight_total_ounces"`
	WeightTotalGrams        float64   `json:"weight_total_grams"`
	CaloriesPerServing      float64   `json:"calories_per_serving"`
	CaloriesTotal           float64   `json:"calories_total"`
	FatsPerServing          float64   `json:"fats_per_serving"`
	FatsTotal               float64   `json:"fats_total"`
	CarbohydratesPerServing float64   `json:"carbohydrates_per_serving"`
	CarbohydratesTotal      float64   `json:"carbohydrates_total"`
	FiberPerServing         float64   `json:"fiber_per_serving"`
	FiberTotal              float64   `json:"fiber_total"`
	ProteinPerServing       float64   `json:"protein_per_serving"`
	ProteinTotal            float64   `json:"protein_total"`
	CostPerServing          float64   `json:"cost_per_serving"`
	CostTotal               float64   `json:"cost_total"`
	CreatedAt               time.Time `db:"created_at" json:"created_at"`
	UpdatedAt               time.Time `db:"updated_at" json:"updated_at"`
}

func NewIngredient(name string, brand string, servings float64, packaging []string, weight float64, caloriesPerServing float64, fatsPerServing float64, carbohydratesPerServing float64, fiberPerServing float64, proteinPerServing float64, costTotal float64) *Ingredient {

	return &Ingredient{
		Name:                    name,
		Brand:                   brand,
		Servings:                servings,
		Packaging:               packaging,
		WeightTotalOunces:       0,
		WeightTotalGrams:        0,
		CaloriesPerServing:      caloriesPerServing,
		CaloriesTotal:           caloriesPerServing * servings,
		FatsPerServing:          fatsPerServing,
		FatsTotal:               fatsPerServing * servings,
		CarbohydratesPerServing: carbohydratesPerServing,
		CarbohydratesTotal:      carbohydratesPerServing * servings,
		FiberPerServing:         fiberPerServing,
		FiberTotal:              fiberPerServing * servings,
		ProteinPerServing:       proteinPerServing,
		ProteinTotal:            proteinPerServing * servings,
		CostPerServing:          math.Round(costTotal/servings*100) / 100,
		CostTotal:               costTotal,
	}
}

func CalcOuncesFromGrams(grams float64) float64 {
	return math.Round((grams/28.3495)*100) / 100
}

func CalcGramsFromOunces(ounces float64) float64 {
	return math.Round((ounces*28.3495)*100) / 100
}
