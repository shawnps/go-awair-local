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
	tempStr := fmt.Sprintf("%.2f", latestData.Temp)
	humidStr := fmt.Sprintf("%.2f", latestData.Humid)
	co2Str := strconv.Itoa(latestData.CO2)
	vocStr := strconv.Itoa(latestData.VOC)
	pm25Str := strconv.Itoa(latestData.PM25)

	fmt.Fprintf(w, "Score\t%s\n", colorPrintFn("Score", latestData.Score)(scoreStr))
	fmt.Fprintf(w, "Temp\t%s\n", colorPrintFn("Temp", latestData.Temp)(tempStr))
	fmt.Fprintf(w, "Humidity\t%s\n", colorPrintFn("Humid", latestData.Humid)(humidStr))
	fmt.Fprintf(w, "CO2\t%s\n", colorPrintFn("CO2", latestData.CO2)(co2Str))
	fmt.Fprintf(w, "VOC\t%s\n", colorPrintFn("VOC", latestData.VOC)(vocStr))
	fmt.Fprintf(w, "PM2.5\t%s\n", colorPrintFn("PM25", latestData.PM25)(pm25Str))
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
	case "Temp":
		v := value.(float64)
		switch {
		case v >= 18 && v <= 25:
			return green
		case v >= 25 && v <= 26:
			return yellow
		case v >= 17 && v <= 18:
			return yellow
		case v >= 9 && v <= 17:
			return magenta
		case v >= 26 && v <= 34:
			return magenta
		default:
			return red
		}
	case "Humid":
		v := value.(float64)
		switch {
		case v >= 40 && v <= 50:
			return green
		case v >= 35 && v <= 40:
			return yellow
		case v >= 50 && v <= 60:
			return yellow
		case v >= 60 && v <= 80:
			return magenta
		case v >= 15 && v <= 35:
			return magenta
		default:
			return red
		}
	case "CO2":
		v := value.(int)
		switch {
		case v <= 600:
			return green
		case v <= 1000:
			return yellow
		case v <= 1500:
			return magenta
		default:
			return red
		}
	case "VOC":
		v := value.(int)
		switch {
		case v >= 0 && v <= 333:
			return green
		case v >= 333 && v <= 1000:
			return yellow
		case v >= 1000 && v <= 8332:
			return magenta
		default:
			return red
		}
	case "PM25":
		v := value.(int)
		switch {
		case v >= 0 && v <= 15:
			return green
		case v >= 15 && v <= 35:
			return yellow
		case v >= 35 && v <= 75:
			return magenta
		default:
			return red
		}
	}

	return color.New(color.FgYellow).SprintFunc()
}
