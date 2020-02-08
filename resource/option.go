package resource

import "github.com/shriharishastry/brightpearl-golang/connector"

type OptionResource struct {
	resourceUrl string
	connection connector.HttpClient
}

type Option struct{
	OptionName string `json:"product_type_name"`
}

