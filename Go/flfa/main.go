package main

import (
	"fmt"
)

func main() {
	var str string
	state := 0

	fmt.Println("Enter a string")
	fmt.Scan(&str)

	for _, char := range str {
		if (char != '0') && (char != '1') {
			panicMessage := "The character '" + string(char) + "' is not valid"
			panic(panicMessage)
		}

		switch state {
		case 0:
			switch char {
			case '0':
				state = 0
			case '1':
				state = 1
			}
		case 1:
			switch char {
			case '0':
				state = 2
			case '1':
				state = 0
			}
		case 2:
			switch char {
			case '0':
				state = 1
			case '1':
				state = 2
			}
		}
	}

	fmt.Printf("State: %v\n", state)
}
