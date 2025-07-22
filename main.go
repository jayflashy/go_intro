package main

import (
	"fmt"
	"jayflashy/go_intro/prices"

	"rsc.io/quote"
)

func main() {

	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	pastaRecipie := PrintRecipe("pasta")
	fmt.Println(pastaRecipie)

	saladRecipe := PrintRecipe("salad")
	fmt.Println(saladRecipe)

	soupPrice := prices.GetPrice("soup")
	fmt.Println(soupPrice)
}
