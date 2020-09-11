package main

import (
	//Import standard libraries
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"

	//Import third party library
	"github.com/PuerkitoBio/goquery"
)

func Scrape() {

	var N int

	flag.IntVar(&N, "N", 0, "Number of requests")
	flag.Parse()

	for i := 1; i <= N; i++ {
		// Make HTTP request
		res, err := http.Get("https://spb.cian.ru/rent/flat/240283825/")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal("Error loading HTTP response body.", err)
		}

		var metaDescription string
		//Read titile page
		doc.Find("meta").Each(func(index int, item *goquery.Selection) {
			if item.AttrOr("name", "") == "description" {
				metaDescription = item.AttrOr("content", "")
			}
		})
		//Use regular expression to find apartaments for rent
		re := regexp.MustCompile("Цена аренды")
		words := re.FindAllString(string(metaDescription), -1)
		if words == nil {
			fmt.Println("")
		} else {
			fmt.Println(metaDescription)
		}

	}

}
