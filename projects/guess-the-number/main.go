// Guess the Number:

// A game where the program randomly selects a number between 1 and 100, and the user has to guess it. The program provides hints if the guess is too high or too low.
// Key Concepts: Random number generation, loops, user input.

package main

import (
	"fmt"
	"math/rand/v2"
)

var score int = 0

func main() {
	var n int
	finish := "1"
	for finish == "1" {
		randomNumber := randRange(0, 10)
		fmt.Println("Adivine el numero")
		gameLoop(n, randomNumber)
		fmt.Printf("Correcto! El numero era el %d \n", randomNumber)
		fmt.Printf("Su puntaje es de: %d \n", score)
		fmt.Println("Seguir jugando?")
		fmt.Println("1. Si\n2. No")
		fmt.Scan(&finish)
	}
}

func gameLoop(n int, randomNumber int) {
	for n != randomNumber {

		fmt.Scan(&n)
		checkNumber(n, randomNumber)
	}
	score++
}
func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func checkNumber(n int, randomNumber int) {
	if n < randomNumber {
		fmt.Println("El numero ha sido muy bajo")
	}
	if n > randomNumber {
		fmt.Println("El numero ha sido muy alto")
	}

	fmt.Println("Segui intendanto!")
}
