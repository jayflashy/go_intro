package main

import (
	"errors"
	"fmt"
)

func PrintRecipe(food string) (string, error) {
	if food == "" {
		return "", errors.New("no food item was provided")
	}

	title := fmt.Sprintf("Here is a recipe for %s:\n", food)
	switch food {
	case "pasta":
		ingredient := "Ingredients: pasta, water, salt"
		instructions := "Instructions: Boil water, add salt, cook pasta until al dente, drain and serve."
		message := title + ingredient + "\n" + instructions
		return message, nil
	case "salad":
		ingredient := "Ingredients: lettuce, tomatoes, cucumbers, dressing"
		instructions := "Instructions: Chop vegetables, mix with dressing, serve chilled."
		extra := fmt.Sprintf("Enjoy your %s!", food)
		message := title + ingredient + "\n" + instructions + "\n" + extra
		return message, nil
	case "soup":
		ingredient := "Ingredients: broth, vegetables, noodles"
		instructions := "Instructions: Heat broth, add vegetables and noodles, simmer until cooked."
		extra := fmt.Sprintf("Voila! Your %s is ready", food)
		message := title + ingredient + "\n" + instructions + "\n" + extra
		return message, nil
	default:
		return "", fmt.Errorf("sorry, no recipe found for %s", food)
	}
}
