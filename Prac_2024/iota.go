package main

import "fmt"

func main() {

	const (
		a = iota
		b = iota
		_
		_

		c
		d
	)

	fmt.Println(a, b, c, d)

	type Genre int

	// instead of this

	// const (
	// 	Adventure     Genre = 1
	// 	Comic         Genre = 2
	// 	Crime         Genre = 3
	// 	Fiction       Genre = 4
	// 	Fantasy       Genre = 5
	// 	Historical    Genre = 6
	// 	Horror        Genre = 7
	// 	Magic         Genre = 8
	// 	Mystery       Genre = 9
	// 	Philosophical Genre = 10
	// 	Political     Genre = 11
	// 	Romance       Genre = 12
	// 	Science       Genre = 13
	// 	Superhero     Genre = 14
	// 	Thriller      Genre = 15
	// 	Western       Genre = 16
	// )

	// we can use this

	const (
		Adventure Genre = iota
		Comic
		Crime
		Fiction
		Fantasy
		Historical
		Horror
		Magic
		Mystery
		Philosophical
		Political
		Romance
		Science
		Superhero
		Thriller
		Western
	)

	fmt.Println(Adventure, Western, Superhero, Philosophical)
}
