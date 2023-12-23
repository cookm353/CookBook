package main

import (
	"fmt"
	"reflect"

	"github.com/cookm353/CookBook/models"
)

func main() {
	recipe := models.GetRecipeByID("1")

	// for _, recipe := range recipes {
	fmt.Println("Recipe:", recipe)
	v := reflect.ValueOf(recipe)
	typeOfRecipe := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s:\t: %v\n", typeOfRecipe.Field(i).Name, v.Field(i).Interface())
	}

	// }
}
