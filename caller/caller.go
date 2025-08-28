package main

import(
	"fmt"
	"PJ2005/greetings"
)

func main(){
	// get message from greetings module and print
	message := greetings.Hello("Rohan")
	fmt.Println(message)
}