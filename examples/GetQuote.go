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
	fmt.Println(data[0].Quote)
}

// Prints a random quote, i.e: "I used to be with it. But then they changed what it was. Now what I'm with isn't it, and what's it seems scary and wierd. It'll happen to you."
