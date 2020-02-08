package resource

import (
	"fmt"
	"github.com/shriharishastry/brightpearl/connector"
	"log"
	"net/http"
	"net/url"
)

type SeasonResource struct {
	resourceUrl string
	connection connector.HttpClient
}

type Season struct{
	SeasonName string `json:"season_name"`
}


func (c *SeasonResource) GetAll(){
	rsp, err := c.connection.SendRequest(http.MethodGet, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}


func (c *SeasonResource) Get(){
	rsp, err := c.connection.SendRequest(http.MethodGet, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

func (c *SeasonResource) Create()  {
	rsp, err := c.connection.SendRequest(http.MethodPost, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

func (c *SeasonResource) Update()  {
	rsp, err := c.connection.SendRequest(http.MethodPut, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

func (c *SeasonResource) Delete()  {
	rsp, err := c.connection.SendRequest(http.MethodDelete, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

