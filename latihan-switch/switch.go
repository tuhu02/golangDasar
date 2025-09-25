package main

import (
	"fmt"
)

func main(){
	// angka := 2

	// switch(angka){
	// case 1:
	// 	fmt.Println("Angka 1")
	// case 2:
	// 	fmt.Println("Angka 2")
	// case 3:
	// 	fmt.Println("Angka 3")
	// default:
	// 	fmt.Println("Angka tidak dikenal")
	// }

	hari := "minggu"

	switch(hari){
		case "senin", "selasa", "rabu", "kamis":
			fmt.Println("Sekarang hari aktif")
		case "sabtu", "minggu":
			fmt.Println("Sekarang hari libur")
		default:
			fmt.Println("hari tidak valid") 
	}
}