package main

import (
	"fmt"
	"jayflashy/go_intro/prices"

	"rsc.io/quote"
)

func main() {

	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	pastaRecipie, _ := PrintRecipe("pasta")
	fmt.Println(pastaRecipie)

	saladRecipe, _ := PrintRecipe("salad")
	fmt.Println(saladRecipe)

	soupRecipe, error := PrintRecipe("")
	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println(soupRecipe)
	}

	soupPrice := prices.GetPrice("soup")
	fmt.Println(soupPrice)
}
