package main

import "fmt"

func main() {
	var batas int
	var total int
	var b string

	fmt.Print("Batas : ")
	fmt.Scan(&batas)

	if batas % 2 < 0 {
		fmt.Print("Tidak Boleh Negatif")
	} else {
		a := 1
		for a <= batas*2 {
			if a % 2 == 0 {
				total += a
				if a == batas*2 {
					b = " = "
				} else {
					b = "+"
				}
				fmt.Print(a, b)
			}
			a = a + 1
		}
		fmt.Print(total)
	}
	
}