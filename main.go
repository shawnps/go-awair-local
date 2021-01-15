package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/shawnps/go-awair-local/awair"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "", "your Awair's local HTTP server address")
	flag.Parse()

	if addr == "" {
		log.Fatal("-addr is required")
	}

	client := awair.NewClient(addr)

	latestData, err := client.LatestData()
	if err != nil {
		log.Fatal(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

	scoreStr := strconv.Itoa(latestData.Score)

	fmt.Fprintf(w, "Score\t%s\n", colorPrintFn("Score", latestData.Score)(scoreStr))
	fmt.Fprintf(w, "Temp\t%.2f\n", latestData.Temp)
	fmt.Fprintf(w, "Humidity\t%.2f\n", latestData.Humid)
	fmt.Fprintf(w, "CO2\t%d\n", latestData.CO2)
	fmt.Fprintf(w, "VOC\t%d\n", latestData.VOC)
	fmt.Fprintf(w, "PM2.5\t%d\n", latestData.PM25)
	w.Flush()
}

// https://support.getawair.com/hc/en-us/articles/360039242373-Air-Quality-Factors-Measured-By-Awair-Element#h_9c45cca9-012c-4419-ac2c-d3afeb683351
func colorPrintFn(factor string, value interface{}) func(a ...interface{}) string {
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	switch factor {
	case "Score":
		v := value.(int)
		switch {
		case v >= 90:
			return green
		case v >= 80:
			return yellow
		case v >= 60:
			return magenta
		default:
			return red
		}
	}

	return color.New(color.FgYellow).SprintFunc()
}
