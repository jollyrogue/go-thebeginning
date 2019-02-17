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

	charCount, longStrings := printNames(argsPartial)

	fmt.Printf("Longest Strings: %v\n", longStrings)
	fmt.Printf("Number of Characters in strings: %d\n", charCount)
}

func printNames(nameList []string) (characterCount uint, longestString []string) {
	maxLength := 0

	for index, name := range nameList {
		fmt.Printf("Hello, %s.\t(idx: %d, charcount: %d)\n", name, index, len(name))
		characterCount += uint(len(name))

		if len(name) > maxLength {
			maxLength = len(name)
		}
	}

	for _, name := range nameList {
		if len(name) == maxLength {
			longestString = append(longestString, name)
		}
	}

	return
}
