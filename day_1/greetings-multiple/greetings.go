package greetingsmultiple

import (
	greetings_random "greetings-random"
)

func Multiple(names []string) []string {
	var greetings []string
	for _, name := range names {
		greeting, err := greetings_random.Hello(name)
		if err != nil {
			continue
		}
		greetings = append(greetings, greeting)
	}

	return greetings
}
