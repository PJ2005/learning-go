package main

import(
	"fmt"
	"log"

	"PJ2005/greetings"
)

func main(){

	// command at start of each log message
	log.SetPrefix("greetings: ")

	// 0 as we are not putting any other fiekds in the output message
	log.SetFlags(0)

	// get message and error from greetings module
	message, err := greetings.Hello("")

	// if the error message value is not nil, terminate the code and log the error
	if err != nil{
		log.Fatal(err)
	}

	// else print the message returned
	fmt.Println(message)
}