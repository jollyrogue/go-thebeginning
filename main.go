package main

import (
	"fmt"
	"os"
)

func main() {
	argsAll := os.Args
	argsPartial := os.Args[1:]

	fmt.Printf("All Arguments: (%T) %+v\n", argsAll, argsAll)
	fmt.Printf("Arguments: (%T) %+v\n", argsPartial, argsPartial)

	printNames(argsPartial)
}

func printNames(nameList []string) {
	for index, name := range nameList {
		fmt.Printf("Hello, %s.\t(%d)\n", name, index)
	}
}
