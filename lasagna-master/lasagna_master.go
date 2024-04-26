// Package lasagna provides functions for cooking lasagna.
package lasagna

const (
	// DefaultPreparationTimePerLayer represents the default preparation time per layer.
	DefaultPreparationTimePerLayer = 2
	// NoodleQuantityByLayer represents the quantity of noodles required per layer of lasagna.
	NoodleQuantityByLayer = 50
	// SauceQuantityByLayer represents the quantity of sauce required per layer of lasagna.
	SauceQuantityByLayer = 0.2
	// ServingInRecipe represents the default serving in the recipe.
	ServingInRecipe = 2
)

// PreparationTime calculates the total preparation time for cooking the lasagna.
func PreparationTime(layers []string, preparationTimePerLayer int) int {
	if preparationTimePerLayer <= 0 {
		preparationTimePerLayer = DefaultPreparationTimePerLayer
	}
	return len(layers) * preparationTimePerLayer
}

// Quantities calculates the total quantities of noodles and sauce needed to cook the lasagna.
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += NoodleQuantityByLayer
		} else if layer == "sauce" {
			sauce += SauceQuantityByLayer
		}
	}
	return
}

// AddSecretIngredient adds the last ingredient from the friend's list to the cook's list as the secret ingredient.
func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// ScaleRecipe scales the quantities of ingredients in the lasagna recipe to the desired number of portions.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	copy(scaledQuantities, quantities)
	for i, quantity := range scaledQuantities {
		scaledQuantities[i] = quantity * float64(portions) / 2
	}
	return scaledQuantities
}
