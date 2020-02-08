package resource

import "github.com/shriharishastry/brightpearl/connector"

type CustomFieldResource struct {
	resourceUrl string
	connection connector.HttpClient
}

type CustomField struct{
	CustomFieldName string `json:"product_name"`
}

