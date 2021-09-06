package main

import "fmt"

func main() {
	var batas int

	fmt.Print("Batas : ")
	fmt.Scan(&batas)

	a := 1
	

	if batas % 2 < 0 {
		fmt.Print("Tidak Boleh Negatif")
	} else {
		for a <= batas {
			if a % 2 == 0 {
				println(a+3)
			} else {
				println(a)
			}
			
			a+=1
		}
	}
	
	
}