package main

import (
	"flag"
	"log"
)

var n int

func main() {

	flag.IntVar(&n, "n", 0, "Number of requests")
	flag.Parse()

	err := Scrape(n, "https://lipetsk.cian.ru/rent/flat/241832915/")
	if err != nil {
		log.Fatal(err)
	}

}
