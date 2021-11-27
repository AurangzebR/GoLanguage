package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "name", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}
func init() {
	rand.Seed(time.Now().UnixNano())
}
func randomFormat() string {
	formats := []string{
		"hi, %v. welcome",
		"Great to see you,%v",
		"Hail, %v! well met!"}
	return formats[rand.Intn(len(formats))]

}
