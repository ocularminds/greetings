package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//set initial values
func init() {
	rand.Seed(time.Now().UnixNano())
}

//Hello Prints gtreeting message for the name
//or throws error is name is nill
func Hello(name string) (string, error) {
	//if no name is given return error message
	if name == "" {
		return "", errors.New("Name is required")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

//Hellos return greetings for multiple people
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := HelloRandom(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

//HelloRandom creates greetings with random format
func HelloRandom(name string) (string, error) {
	if name == "" {
		return "", errors.New("name error: We need your name to greet")
	}
	return fmt.Sprintf(randomFormat(), name), nil
}

func randomFormat() string {
	//array with no length is dynamic array
	formats := []string{"Hi, %v. Welcome", "Greet to see you, %v!", "Hail %v. Well met!"}
	return formats[rand.Intn(len(formats))]
}

//Greet greetings hello
func Greet() string {
	return "Hello GitHub Actions"
}
