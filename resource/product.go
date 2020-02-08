package resource

import (
	"fmt"
	"github.com/shriharishastry/brightpearl-golang/connector"
	"log"
	"net/http"
	"net/url"
)

type ProductResource struct {
	resourceUrl string
	connection connector.HttpClient
}

type Product struct{
	ProductName string `json:"product_name"`
}


func (p *ProductResource) GetAll(){
	rsp, err := p.connection.SendRequest(http.MethodGet, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}


func (p *ProductResource) Get(){
	rsp, err := c.connection.SendRequest(http.MethodGet, c.resourceUrl, url.Values{}, http.Header{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

func (p *ProductResource) Create()  {

}

func (p *ProductResource) Update()  {

}

func (p *ProductResource) Delete()  {

}