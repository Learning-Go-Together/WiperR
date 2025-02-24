package learn_errors

import (
	"errors"
)

func DivideWithError(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("No dividing by zero")
	}
	return x / y, nil
}
