package resource

import (
	"fmt"
	"github.com/shriharishastry/brightpearl/connector"
	"log"
	"net/http"
	"net/url"
)

type ProductTypeResource struct {
	resourceUrl string
	connection connector.HttpClient
}

type ProductType struct{
	ProductTypeName string `json:"product_type_name"`
}


func (c *ProductTypeResource) GetAll(){
	rsp, err := c.connection.SendRequest(http.MethodGet, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}


func (c *ProductTypeResource) Get(){
	rsp, err := c.connection.SendRequest(http.MethodGet, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

func (c *ProductTypeResource) Create()  {
	rsp, err := c.connection.SendRequest(http.MethodPost, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

func (c *ProductTypeResource) Update()  {
	rsp, err := c.connection.SendRequest(http.MethodPut, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

func (c *ProductTypeResource) Delete()  {
	rsp, err := c.connection.SendRequest(http.MethodDelete, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

