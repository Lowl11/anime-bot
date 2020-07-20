package parser

import (
	"io/ioutil"
	"log"
	"net/http"
)

func getResponseOfPage(url string) []byte {
	response, err := http.Get(AnistarTestURL)
	if err != nil {
		log.Panic(err)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panic(err)
	}
	return responseBody
}
