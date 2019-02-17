package main

import (
	"fmt"
	"os"
)

func main() {
	argsAll := os.Args
	argsActual := os.Args[1:]

	fmt.Printf("All Arguments: %+v\n", argsAll)
	fmt.Printf("Arguments: %+v\n", argsActual)
	fmt.Println("Hello.")
}
