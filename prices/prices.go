package prices

func GetPrice(recipe string) string {
	switch recipe {
	case "pasta":
		return "The price of pasta is $2.50 per pound."
	case "salad":
		return "The price of salad ingredients is approximately $3.00 per serving."
	case "soup":
		return "The price of soup ingredients is around $4.00 per serving."
	default:
		return "Sorry, I don't have a price for that recipe."
	}
}
