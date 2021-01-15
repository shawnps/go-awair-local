package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/shawnps/go-awair-local/awair"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "", "your Awair's local HTTP server address")
	flag.Parse()

	client := awair.NewClient(addr)

	latestData, err := client.LatestData()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(latestData)

}
