package resource

import (
	"fmt"
	"github.com/shriharishastry/brightpearl/connector"
	"log"
	"net/http"
	"net/url"
)

type BrandtResource struct {
	resourceUrl string
	connection connector.HttpClient
}

type Brand struct{
	BrandName string `json:"product_name"`
}

func (c *BrandtResource) GetAll(){
	rsp, err := c.connection.SendRequest(http.MethodGet, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}


func (c *BrandtResource) Get(){
	rsp, err := c.connection.SendRequest(http.MethodGet, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

func (c *BrandtResource) Create()  {
	rsp, err := c.connection.SendRequest(http.MethodPost, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

func (c *BrandtResource) Update()  {
	rsp, err := c.connection.SendRequest(http.MethodPut, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

func (c *BrandtResource) Delete()  {
	rsp, err := c.connection.SendRequest(http.MethodDelete, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}
