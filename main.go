package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

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

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

	fmt.Fprintf(w, "Score\t%d\n", latestData.Score)
	fmt.Fprintf(w, "Temp\t%f\n", latestData.Temp)
	fmt.Fprintf(w, "Humidity\t%f\n", latestData.Humid)
	fmt.Fprintf(w, "CO2\t%d\n", latestData.CO2)
	fmt.Fprintf(w, "VOC\t%d\n", latestData.VOC)
	fmt.Fprintf(w, "PM2.5\t%d\n", latestData.PM25)
	w.Flush()
}
