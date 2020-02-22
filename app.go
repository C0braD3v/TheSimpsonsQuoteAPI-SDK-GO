package simpsons

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Quote struct {
	Quote              string `json:"quote"`
	Character          string `json:"character"`
	Image              string `json:"image"`
	CharacterDirection string `json:"characterDirection"`
}

func GetQuotes(count string) ([]Quote, error) {
	quotes := []Quote{}
	response, err := http.Get("https://thesimpsonsquoteapi.glitch.me/quotes?count=" + count)
	if err != nil {
		return nil, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	e := json.Unmarshal(responseData, &quotes)
	if e != nil {
		return nil, e
	}
	return quotes, nil
}
