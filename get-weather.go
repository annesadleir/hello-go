package main

import ("fmt"
"net/http"
"io/ioutil"
"os"
//"encoding/xml"
)

const DarkSkyUri = "https://api.darksky.net/forecast/%s/50.7184,-3.5339?exclude=currently,minutely,daily,alerts,flags"

func main() {
	completedUri := makeUri()
	resp, _ := http.Get(completedUri)
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)
	fmt.Println(stringBody)
	resp.Body.Close()
}

func makeUri() string {
	API_KEY := os.Getenv("DARKSKY_KEY")
	fmt.Println(API_KEY)
	return fmt.Sprintf(DarkSkyUri, API_KEY)
}