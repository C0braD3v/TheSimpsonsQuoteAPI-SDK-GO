package main

import (
	"fmt"
	"log"
	simpsons "simpsons/simpsons"
)

func main() {
	data, err := simpsons.GetQuotes("0")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data[0])
}
