package main

import "fmt"

func main() {
	var s, p int
	fmt.Scan(&s, &p)

	if (s >= 0 && s <= 1000000) && (p >= 0 && p <= 1000000) {
		if s >= p {
			fmt.Print("Pembayaran Berhasil")
		} else {
			fmt.Print("Saldo Tidak Cukup")
		}
	}

}
