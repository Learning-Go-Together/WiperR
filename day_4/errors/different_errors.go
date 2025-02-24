package learn_errors

import (
	"errors"
	"log"
)

func validateStatus(status string) error {
	if len(status) == 0 {
		return errors.New("status can't be empty")
	}

	if len(status) > 140 {
		return errors.New("status exceeds 140 characters")
	}
	return nil
}

func DifferentErrors() {
	err := validateStatus("Hello")
	if err != nil {
		log.Fatal(err)
	}
	err = validateStatus("")
	if err != nil {
		log.Fatal(err)
	}
}
