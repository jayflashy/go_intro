package main

import "fmt"

func PrintRecipe(food string) string {
	title := fmt.Sprintf("Here is a recipe for %s:\n", food)
	switch food {
	case "pasta":
		ingredient := "Ingredients: pasta, water, salt"
		instructions := ("Instructions: Boil water, add salt, cook pasta until al dente, drain and serve.")
		return title + ingredient + "\n" + instructions
	case "salad":
		ingredient := "Ingredients: lettuce, tomatoes, cucumbers, dressing"
		instructions := "Instructions: Chop vegetables, mix with dressing, serve chilled."
		extra := fmt.Sprintf("Enjoy your %s!", food)
		return title + ingredient + "\n" + instructions + "\n" + extra
	case "soup":
		ingredient := "Ingredients: broth, vegetables, noodles"
		instructions := "Instructions: Heat broth, add vegetables and noodles, simmer until cooked."
		extra := fmt.Sprintf("Volaa your %s! is ready", food)
		return title + ingredient + "\n" + instructions + "\n" + extra
	default:
		return fmt.Sprintf("Sorry, I don't have a recipe for %s", food)
	}
}
