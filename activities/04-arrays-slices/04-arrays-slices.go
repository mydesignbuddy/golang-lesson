package main

import "fmt"

func main() {
	// arrays
	var favoriteColors [3]string

	favoriteColors[0] = "blue"
	favoriteColors[1] = "red"
	favoriteColors[2] = "green"
	fmt.Println(favoriteColors)

	// setting an literal array
	favoriteSongs := [2]string{"Every Rose Has Its Thorn", "Livin' La Vida Loca"}
	fmt.Println(favoriteSongs)

	// literal slices
	favoriteMovies := []string{"Star Wars", "Matrix"}

	// append two more elements
	favoriteMovies = append(favoriteMovies, "The Notebook")
	favoriteMovies = append(favoriteMovies, "Home Alone")

	fmt.Println(favoriteMovies)

	fmt.Println(len(favoriteMovies)) // count of elements in slice
}
