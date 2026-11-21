package main

import "fmt"

func main() {
	var a, b, durasi int
	fmt.Scan(&a, &b)

	if b >= a {
		durasi = b - a
	} else {
		durasi = 24 - a + b
	}

	fmt.Println(durasi)
}