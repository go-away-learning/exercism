// Package thefarm provides functionality for calculating and dividing food
// for cows.
package thefarm

import "fmt"

// InvalidCowsError represents an error when the number of cows is invalid.
type InvalidCowsError struct {
	numberOfCows int
	message      string
}

// Error returns a formatted error message for InvalidCowsError.
func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.numberOfCows, err.message)
}

// DivideFood calculates the amount of food per cow based on the fodder amount and fattening factor.
func DivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
	fodderAmount, err := fc.FodderAmount(numberOfCows)
	if err != nil {
		return 0.0, err
	}
	fatteningFactor, err := fc.FatteningFactor()
	if err != nil {
		return 0.0, err
	}
	foodPerCow := (fodderAmount * fatteningFactor) / float64(numberOfCows)
	return foodPerCow, nil
}

// ValidateInputAndDivideFood validates the number of cows and calculates the amount of food per cow if valid.
func ValidateInputAndDivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows <= 0 {
		return 0.0, fmt.Errorf("invalid number of cows")
	}
	return DivideFood(fc, numberOfCows)
}

// ValidateNumberOfCows checks if the number of cows is valid and returns an appropriate error if not.
func ValidateNumberOfCows(numberOfCows int) error {
	if numberOfCows == 0 {
		return &InvalidCowsError{
			numberOfCows: numberOfCows,
			message:      "no cows don't need food",
		}
	} else if numberOfCows < 0 {
		return &InvalidCowsError{
			numberOfCows: numberOfCows,
			message:      "there are no negative cows",
		}
	}
	return nil
}
