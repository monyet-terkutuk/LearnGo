package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println(counter)
		counter = counter + 1
		fmt.Println(counter)
	}

}
