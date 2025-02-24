package learn_errors

import "fmt"

type DivideError struct {
	divident float64
}

func (d DivideError) Error() string {
	return fmt.Sprintf("Cannot divide %v by zero", d.divident)
}

func divide(divident, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, DivideError{divident: divident}
	}

	return divident / divisor, nil
}

func ErrorInterfaces() {
	var ok error
	var res float64

	res, ok = divide(500, 10)
	if ok != nil {
		fmt.Printf("An error encountered: %v\n", ok)
	} else {
		fmt.Printf("%.2f\n", res)
	}

	res, ok = divide(500, 0)
	if ok != nil {
		fmt.Printf("An error encountered: %v\n", ok)
	} else {
		fmt.Printf("%.2f\n", res)
	}
}
