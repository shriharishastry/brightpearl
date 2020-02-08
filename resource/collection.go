package resource


import "github.com/shriharishastry/brightpearl/connector"

type CollectionResource struct {
	resourceUrl string
	connection connector.HttpClient
}

type Collection struct{
	CollectionName string `json:"product_name"`
}


