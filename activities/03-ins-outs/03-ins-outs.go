package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Write to the console
	fmt.Println("Hello, NOW GO!")

	// declare a input scanner to start capturing user input
	scanner := bufio.NewScanner(os.Stdin)

	// loop
	// setup for loop for new inputs
	for scanner.Scan() {
		text := scanner.Text()

		// conditional
		// If user enters "quit"
		//   Print the message "Please Go!"
		//   break out of the for loop
		if text == "quit" {
			fmt.Println("Good day to you!")
			break
		}

		// Print the message "Please Go!"
		fmt.Println("Please Go!")
	}
}
