// GoDo is a to do app that keeps track of items in a list.
//
// * When you launch the app it should print a default list of items
// * If you want to add a new item just type the name and press `enter`
// * If you want to remove an existing item just re-type an item already
//   in the list and press `enter`
// * If you wish to exit the app just type `quit` and press `enter`
//
// This file has two functions to help them remove and get the
// index number of an element in a slice. You will need to use these
// functions to finish the activity.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Declare default list of items
	items := []string{"milk", "eggs", "bread"}

	// Print app introduction message
	fmt.Printf("GoDo - The best TODO app written in Go\n")
	// Print current list to user
	renderList(items)

	// Declare a bufio Scanner to read user inputs
	scanner := bufio.NewScanner(os.Stdin)

	// create a loop and each time the scanner scans
	for scanner.Scan() {
		//// created a variable that stores the
		//// users latest input as a string
		text := scanner.Text()

		//// if the user's input equals
		//// "quit" break out of the loop
		if text == "quit" {
			fmt.Println("goodbye")
			break
			//// else if user's input equals an existing item
			//// remove it from the item from the slice
		} else if indexOf(items, text) > -1 {
			items = removeItem(items, text)
			//// else if the user's input is a new item append
			//// it to the slice
		} else {
			items = append(items, text)
		}

		//// re-render the list of items
		renderList(items)
	}
}

// Create a function that loops thru items and prints the list
func renderList(items []string) {
	fmt.Println("=============================")
	for i := range items {
		fmt.Printf("%d %s\n", i+1, items[i])
	}
}

// GoLang does not have a function to remove an element
// from a slice. This helper function returns a new slice
// with a given item removed.
func removeItem(items []string, value string) []string {
	index := indexOf(items, value)
	if index > -1 {
		return append(items[:index], items[index+1:]...)
	}
	return items
}

// GoLang does not have a function to find the index
// an element in an slice. This helper function that loops
// thru and returns the index of an item in a slice
func indexOf(items []string, value string) int {
	for i := range items {
		if items[i] == value {
			return i
		}
	}
	return -1
}
