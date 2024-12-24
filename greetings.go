package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
// If the provided name is an empty string, it returns an error indicating that the name is empty.
// Otherwise, it formats a welcome message including the provided name and returns it.
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		if message, err := Hello(name); err != nil {
			return nil, err
		} else {
			messages[name] = message
		}
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, welcome %v!\n",
		"Great to see you %v!\n",
		"Hail, %v! Well met!\n",
	}
	return formats[rand.Intn(len(formats))]
}
