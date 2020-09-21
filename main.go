package main

import "flag"

var N int

func main() {

	flag.IntVar(&N, "N", 0, "Number of requests")
	flag.Parse()
	Scrape(N)

}
