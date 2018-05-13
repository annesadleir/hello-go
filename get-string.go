package main

import ("fmt"
"net/http"
"io/ioutil"
"encoding/xml")

func main() {
	resp, _ := http.Get("http://feeds.bbci.co.uk/news/technology/rss.xml?edition=uk")
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)
	fmt.Println(stringBody)
	resp.Body.Close()

}