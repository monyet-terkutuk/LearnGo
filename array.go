package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "Gabriel"
	name[1] = "Yonathan"
	name[2] = "Siagian"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	// Membuat Array sekaligus membuat isi nya
	var haris = [4]int{
		32,
		34,
		43,
		34,
	}

	fmt.Println(haris)
	fmt.Println(name)
}
