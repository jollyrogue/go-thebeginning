package main

import (
	"fmt"
	"os"
)

func main() {
	argsAll := os.Args
	argsPartial := os.Args[1:]

	fmt.Printf("All Arguments: %+v\n", argsAll)
	fmt.Printf("Arguments: %+v\n", argsPartial)

	for index, value := range argsPartial {
		fmt.Printf("Hello, %s.\t(%d)\n", value, index)
	}
}
