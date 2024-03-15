package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// initialize sqlx pool
	// _, err = NewPGPool()
	// if err != nil {
	// log.Fatal(err)
	// }

	beans := NewIngredient("beans", "stater brothers", 3.5, []string{"can"}, 454, 110, 2.5, 16, 7, 6, 1.39)
	tortillas := NewIngredient("tortills", "stater brothers", 3.5, []string{"bag"}, 1000, 66, 2, 10, 0, 1, 2.48)
	cheese := NewIngredient("cheese", "stater brothers", 3.5, []string{"bag"}, 454, 140, 9, 1, 0, 8, 5.50)

	tacos := NewRecipe("Tacos", map[*Ingredient]float64{beans: 1, tortillas: 0.4, cheese: 0.5})

	tacos.Display()
}
