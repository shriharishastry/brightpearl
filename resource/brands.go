package resource

import "github.com/shriharishastry/brightpearl-golang/connector"

type BrandtResource struct {
	resourceUrl string
	connection connector.HttpClient
}

type Brand struct{
	BrandName string `json:"product_name"`
}

