package resource

import "github.com/shriharishastry/brightpearl-golang/connector"

type ProductTypeResource struct {
	resourceUrl string
	connection connector.HttpClient
}

type ProductType struct{
	ProductTypeName string `json:"product_type_name"`
}

