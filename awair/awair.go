package awair

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// AirData contains the data returned by the Awair local HTTP server's
// /air-data/latest endpoint
type AirData struct {
	Timestamp     time.Time
	Score         int
	DewPoint      float64 `json:"dew_point"`
	Temp          float64
	Humid         float64
	AbsHumid      float64 `json:"abs_humid"`
	CO2           int
	CO2Est        int `json:"co2_est"`
	VOC           int
	VOCBaseline   int `json:"voc_baseline"`
	VOCH2Raw      int `json:"voc_h2_raw"`
	VOCEthanolRaw int `json:"voc_ethanol_raw"`
	PM25          int
	PM10Est       int `json:"pm10_est"`
}

// Client contains the Awair's LAN IP for making requests to the
// local API
type Client struct {
	Addr string
}

// NewClient returns a new Awair local client
func NewClient(addr string) *Client {
	return &Client{
		Addr: addr,
	}
}

// LatestData gets the latest data from the Awair local HTTP server
func (c *Client) LatestData() (AirData, error) {
	var ad AirData
	resp, err := http.Get(fmt.Sprintf("http://%s/air-data/latest", c.Addr))
	if err != nil {
		return ad, err
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	err = dec.Decode(&ad)
	return ad, err
}
