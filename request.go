package main

import (
	//Import standard libraries

	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	//Import third party library
	"github.com/PuerkitoBio/goquery"
)

func Scrape(N int) (string, error) {
	for i := 1; i <= N; i++ {
		// Make HTTP request
		res, err := http.Get("https://www.cian.ru/rent/flat/235640722/")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal("Error loading HTTP response body.", err)
		}

		var metaDescription string
		var price string
		var address string
		var format string
		//Head description
		doc.Find("meta").Each(func(index int, item *goquery.Selection) {
			if item.AttrOr("name", "") == "description" {
				metaDescription = item.AttrOr("content", "")
			}
		})

		doc.Find("span").Each(func(index int, item *goquery.Selection) {
			if item.AttrOr("itemprop", "") == "name" {
				address = item.AttrOr("content", "")
			}
		})

		doc.Find("h1").Each(func(i int, item *goquery.Selection) {
			item.SetAttr("class", "title")
			format, _ = item.Attr("class")
			if format == "title" {
				format = item.Text()
			}
		})

		doc.Find("span").Each(func(index int, item *goquery.Selection) {
			if item.AttrOr("itemprop", "") == "price" {
				price = item.AttrOr("content", "")
			}
		})

		//Use regular expression to find apartaments for rent
		re := regexp.MustCompile("Цена аренды")
		words := re.FindAllString(string(metaDescription), -1)
		if words == nil {
			fmt.Println("")
		} else {
			fmt.Println("Адрес:", address, " Формат:", format, " Цена:", price)
		}

	}
	return "", errors.New("this is an error")

}
