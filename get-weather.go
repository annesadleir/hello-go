package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"encoding/json"
)

const DarkSkyUri = "https://api.darksky.net/forecast/%s/50.7184,-3.5339?exclude=currently,minutely,daily,alerts,flags"

func main() {
	completedUri := makeUri()
	resp, _ := http.Get(completedUri)
	bytes, _ := ioutil.ReadAll(resp.Body)
	var wd Weather
	err := json.Unmarshal(bytes, &wd)
	fmt.Println(err)
	fmt.Println(wd.Latitude)
	resp.Body.Close()
}

func makeUri() string {
	API_KEY := os.Getenv("DARKSKY_KEY")
	fmt.Println(API_KEY)
	return fmt.Sprintf(DarkSkyUri, API_KEY)
}
