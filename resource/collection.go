package resource

import (
	"fmt"
	"github.com/shriharishastry/brightpearl/connector"
	"log"
	"net/http"
	"net/url"
)

type CollectionResource struct {
	resourceUrl string
	connection connector.HttpClient
}

type Collection struct{
	CollectionName string `json:"product_name"`
}


func (c *CollectionResource) Get(){
	rsp, err := c.connection.SendRequest(http.MethodGet, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

func (c *CollectionResource) Create()  {
	rsp, err := c.connection.SendRequest(http.MethodPost, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

func (c *CollectionResource) Update()  {
	rsp, err := c.connection.SendRequest(http.MethodPut, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

func (c *CollectionResource) Delete()  {
	rsp, err := c.connection.SendRequest(http.MethodDelete, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}
