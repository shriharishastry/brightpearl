package resource


import "github.com/shriharishastry/brightpearl-golang/connector"

type CollectionResource struct {
	resourceUrl string
	connection connector.HttpClient
}

type Collection struct{
	CollectionName string `json:"product_name"`
}


