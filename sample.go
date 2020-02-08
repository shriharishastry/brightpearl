package main

import (
	"fmt"
	"github.com/shriharishastry/brightpearl-golang/connector"
	"net/url"
)

func main() {
	client := connector.HttpClient{}
	params:= url.Values{}
	params.Add("q", "title:DNA")
	resp, err := client.SendRequest("GET", "http://api.plos.org/", params, nil)
	fmt.Println(resp, err)

}
