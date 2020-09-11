package main

import (
	"fmt"
	"log"

	"ocularminds.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0) //disbale date and time
	message, err := greetings.Hello("Festus")
	fmt.Println(message)

	message, err = greetings.HelloRandom("Daniels")
	fmt.Println(message)
	names := []string{"Gladys", "Bunmi", "Fatmat", "Ngozi"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

	message, err = greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

}
