package main // All functions from files in same directory

import (
	"fmt" // Formatting text and printing to console. Standard with Go install

	"rsc.io/quote/v4" // Imported external package
)

func main() {
	fmt.Println(quote.Go()) // print a random Go quote using an external quote package

}
