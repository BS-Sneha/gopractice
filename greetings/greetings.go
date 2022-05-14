package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	//msg := fmt.Sprintf(randomFormat(), name)
	msg := fmt.Sprint(randomFormat())
	return msg, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
func Hellos(names []string) (map[string]string, error) {
	message := make(map[string]string)
	for _, name := range names {
		messages, err := Hello(name)
		if err != nil {
			return nil, err
		}
		message[name] = messages
	}
	return message, nil
}
