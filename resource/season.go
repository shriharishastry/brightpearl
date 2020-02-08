package resource

import "github.com/shriharishastry/brightpearl/connector"

type SeasonResource struct {
	resourceUrl string
	connection connector.HttpClient
}

type Season struct{
	SeasonName string `json:"season_name"`
}

