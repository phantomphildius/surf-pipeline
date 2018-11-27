package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type swell struct {
	Components struct {
		Combined struct {
			Direction string `json:"compassDirection"`
			Height    int
			Period    int
		}
	}
}
type wind struct {
	Speed     float64 `json:"speed"`
	Direction string  `json:"compassDirection"`
	Chill     int     `json:"chill"`
}
type weather struct {
	Temp int `json:"temperature"`
	Wind wind
}
type forecast struct {
	Rating  int `json:"solidRating"`
	Time    int `json:"localTimestamp"`
	Wave    swell
	Weather weather
}

func main() {
	url := fmt.Sprintf("https://magicseaweed.com/api/%v/forecast?spot_id=846&units=us", os.Getenv("MAGIC_SEAWEED_API_KEY"))
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fs := &[]forecast{}
	e := json.Unmarshal(bs, fs)
	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(1)
	}

	fmt.Println(fs)
}
