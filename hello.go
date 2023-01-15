package main

import "fmt"

func main() {
	var metin string = "Merhaba Dunya"
	var kdv int = 18
	fmt.Println("Hello world")
	fmt.Println(metin)
	fmt.Println(100 + (100 * kdv / 100))

	var kdv2 float32 = 0.1
	fmt.Println(100 + 100*kdv2)

	kdv3 := 25
	fmt.Println(kdv3)
	fmt.Printf("veri tipi: %T", kdv3)
}
