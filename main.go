package main

import (
	"flag"
	"fmt"
	"os"
)

var printLongestNames bool

func init() {
	flag.BoolVar(&printLongestNames, "l", false, "Prints the longest names passed in the arguments.")
}

func main() {
	argsAll := os.Args
	argsPartial := os.Args[1:]

	flag.Parse()

	fmt.Printf("All Arguments: (%T) %+v\n", argsAll, argsAll)
	fmt.Printf("Arguments: (%T) %+v\n", argsPartial, argsPartial)

	charCount, longStrings := printNames(argsPartial)

	if printLongestNames == true {
		fmt.Printf("Longest Strings: %v\n", longStrings)
	} else {
		fmt.Println("\033[38;5;196m>>>\033[0m    \033[38;5;21;48;5;118mPrinting of longest names has been turned off!\033[0m    \033[38;5;196m<<<\033[0m")
	}
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
