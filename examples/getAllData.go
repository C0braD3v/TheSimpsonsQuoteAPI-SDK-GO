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

// Prints all data from the random quote. i.e: {Facts are meaningless. You could use facts to prove anything that's even remotely true.  https://cdn.glitch.com/3c3ffadc-3406-4440-bb95-d40ec8fcde72%2FHomerSimpson.png?1497567511939 }
