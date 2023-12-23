package models

import (
	"fmt"
)

type Recipe struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description:`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
	Unit   string `json:"unit"`
}

func NewRecipe(id int, name, description string, ingredients []Ingredient) *Recipe {
	recipe := Recipe{
		ID:          id,
		Name:        name,
		Description: description,
		Ingredients: ingredients,
	}
	return &recipe
}

func (r Recipe) GetAllRecipes() *[]Recipe {
	db := ConnectToDB()
	defer db.Close()

	var recipes *[]Recipe

	return recipes
}

// func (r *Recipe) GetRecipeByID(id string) Recipe {
func GetRecipeByID(id string) Recipe {
	db := ConnectToDB()
	defer db.Close()
	var recipe Recipe

	row, err := db.Query("SELECT * FROM recipes WHERE id = ?", id)
	if err != nil {
		fmt.Println("Row not found")
		return recipe
	}

	fmt.Println("Row:", row)

	recipe = *NewRecipe(
		2,
		"Ragu",
		"A rich meat sauce for use with pasta",
		nil,
	)
	return recipe
}

func AddRecipe() Recipe {
	db := ConnectToDB()
	defer db.Close()

	var recipe Recipe

	return recipe
}

func UpdateRecipe() Recipe {
	db := ConnectToDB()
	defer db.Close()

	var recipe Recipe

	return recipe
}

func DeleteRecipe() {
	db := ConnectToDB()
	defer db.Close()
}
