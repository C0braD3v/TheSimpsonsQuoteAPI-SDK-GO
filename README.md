# TheSimpsonsQuoteAPI-SDK-GO
A GO SDK for the SimpsonsQuote API by JLuboff. (https://github.com/JLuboff/TheSimpsonsQuoteAPI)

# Usage

`go get https://github.com/C0braD3v/TheSimpsonsQuoteAPI-SDK-GO/src`

```go
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
```

Data methods:

```
.Quote // The random quote that was fetched
.Image // The image of the character who said the quote
.Character // The character who said the quote
.CharacterDirection // The direction the character was facing```
