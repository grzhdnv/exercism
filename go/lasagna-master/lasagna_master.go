package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2
	}
	return len(layers) * prepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {

	for _, l := range layers {
		switch l {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := make([]float64, len(quantities))
	for i, q := range quantities {
		scaled[i] = q * float64(portions) / 2
	}

	return scaled
}
