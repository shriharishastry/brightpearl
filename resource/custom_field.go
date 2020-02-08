package resource

import "github.com/shriharishastry/brightpearl-golang/connector"

type CustomFieldResource struct {
	resourceUrl string
	connection connector.HttpClient
}

type CustomField struct{
	CustomFieldName string `json:"product_name"`
}

