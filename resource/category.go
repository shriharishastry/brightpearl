package resource

import (
	"fmt"
	"github.com/shriharishastry/brightpearl/connector"
	"log"
	"net/http"
	"net/url"
)

type CategoryResource struct {
	resourceUrl string "category"
	connection connector.HttpClient
}

type Category struct{
	CategoryName string `json:"category_name"`
}


func (c *CategoryResource) GetAll(){
	categories := []Category{}
	fmt.Println(categories)
	rsp, err := c.connection.SendRequest(http.MethodGet, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}


func (c *CategoryResource) Get(){
	rsp, err := c.connection.SendRequest(http.MethodGet, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

func (c *CategoryResource) Create()  {
	rsp, err := c.connection.SendRequest(http.MethodPost, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

func (c *CategoryResource) Update()  {
	rsp, err := c.connection.SendRequest(http.MethodPut, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

func (c *CategoryResource) Delete()  {
	rsp, err := c.connection.SendRequest(http.MethodDelete, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}
