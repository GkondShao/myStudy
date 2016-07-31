package main

import "fmt"

func main() {
	var i int

	for {
		if i == 5 {
			return
		}

		i++
	}

	fmt.Println(i)
}
