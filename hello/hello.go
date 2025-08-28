package main

// formatting package
import "fmt"

// external package for quote
import "rsc.io/quote"

// exected by default when main package is run
func main() {
    fmt.Println(quote.Go())
}