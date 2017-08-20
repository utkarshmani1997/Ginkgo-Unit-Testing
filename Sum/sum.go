package main

import "fmt"

func Sum(x, y int) int {
	fmt.Printf("Received Digits x: %d, y: %d\n", x, y)
	return x + y
}

func main() {
	m, n := 5, 6
	fmt.Println("Sum of two digits : ", Sum(m, n))

}
