package main

import (
	"log"
	"net/http"
)

func main() {
	// Make HTTP request
	response, err := http.Get("https://www.cian.ru/rent/flat/235640722/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

}
