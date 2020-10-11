package main

import (
	//Import standard libraries

	"log"
	"net/http"
	"regexp"

	//Import third party library
	"github.com/PuerkitoBio/goquery"
)

//Variables for search appartemnts
var metaDescription, price, address, format string

func Scrape(n int, url string) (err error) {

	//Regular expression for compile
	re := regexp.MustCompile("Цена аренды")

	for i := 1; i <= n; i++ {
		//HTTP request

		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return err
		}

		//Head description
		doc.Find("meta").Each(func(index int, item *goquery.Selection) {
			if item.AttrOr("name", "") == "description" {
				metaDescription = item.AttrOr("content", "")
			}
		})
		//Find appartament price address
		doc.Find("span").Each(func(index int, item *goquery.Selection) {
			if item.AttrOr("itemprop", "") == "name" {
				address = item.AttrOr("content", "")
			}
			if item.AttrOr("itemprop", "") == "price" {
				price = item.AttrOr("content", "")
			}
		})
		//Find format appartaments
		doc.Find("h1").Each(func(i int, item *goquery.Selection) {
			item.SetAttr("class", "title")
			format, _ = item.Attr("class")
			if format == "title" {
				format = item.Text()
			}
		})

		//Use regular expression to find apartaments for rent
		words := re.FindAllString(string(metaDescription), -1)
		if words != nil {
			log.Println("Адрес:", address, " Формат:", format, " Цена:", price)
		} else {
			log.Fatal(err)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
	return

}
