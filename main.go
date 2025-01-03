package main

import (
	"fmt"
	"log"
)

func main() {
	var val int
	fmt.Println("print number: ")

	_, err := fmt.Scan(&val)
	if err != nil {
		log.Println(err)
	}

	for i := 0; i <= val; i++ {
		fmt.Printf("%d ", i)
	}
}
