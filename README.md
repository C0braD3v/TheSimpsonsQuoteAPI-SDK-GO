# TheSimpsonsQuoteAPI-SDK-GO
An API Wrapper in GoLang for [The SimpsonsQuote API](https://github.com/JLuboff/TheSimpsonsQuoteAPI).

# Usage

`go get https://github.com/HPaulson/Go-Simpsons-Quotes/src`

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
Data Structure:

```go
simpsons {
	GetQuotes("<INT>") // Array<{data}>
}
	
data {
	Quote // String<Quote>
	Image // String<IMG_URL>
	Character // String<Character>
	CharacterDirection // String<Left | Right>
}
```
