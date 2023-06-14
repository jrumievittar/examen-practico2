package main

import "fmt"

func suma(x int, y int) int {
	return x + y
}

func resta(x int, y int) int {
    return x - y
}

func main() {
	fmt.Println(suma(4, 3))
	fmt.Println(resta(8, 2))
}