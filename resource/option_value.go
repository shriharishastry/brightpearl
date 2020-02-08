package resource

import "github.com/shriharishastry/brightpearl-golang/connector"

type OptionValueResource struct {
	resourceUrl string
	connection connector.HttpClient
}

type OptionValue struct{
	OptionValueName string `json:"product_type_name"`
}

