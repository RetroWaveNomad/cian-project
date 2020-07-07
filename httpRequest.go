package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	// Make HTTP request
	response, err := http.Get("https://www.cian.ru/rent/flat/235640722/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Read response data in to memory
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}

	// Create a regular expression to find appartments for rent
	re := regexp.MustCompile("₽/мес.")
	words := re.FindAllString(string(body), -1)
	if words == nil {
		fmt.Println("No matches.")
	} else {
		fmt.Println(" адрес: Москва, улица Пушкина, дом 10; формат: двушка; цена: пять десят тысяч рублей в месяц")
	}
}
