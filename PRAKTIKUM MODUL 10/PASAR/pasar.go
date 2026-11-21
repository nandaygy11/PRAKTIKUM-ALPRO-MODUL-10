package main

import "fmt"

func main() {
	var daging, menu string
	fmt.Scan(&daging, &menu)

	fmt.Println("Daging :", daging)
	if menu == "sop"  {
		fmt.Println("Tambahan : sayur")
	} else if menu == "katsu" {
		fmt.Println("Tambahan : tepung")
	} else if menu == "teriyaki" {
		fmt.Println("Tambahan : bombay")
	} else if menu == "semur" {
		fmt.Println("Tambahan : kecap")
	}
}