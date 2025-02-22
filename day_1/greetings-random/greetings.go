package greetingsrandom

import (
	"errors"
	"fmt"
	"math/rand"
)

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Hello, %v. Nice to meet you!",
		"Greetings, %v. How do you do?",
	}

	return formats[rand.Intn(len(formats))]
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
