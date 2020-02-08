package main

import "time"

type BrightpearlAPI struct {
	ClientId string
	ClientSecret string
	Account string
	Domain string
	AppRef string
	DeveloperRef string

}

const (
	name string = "brightpearlgolangsdk"
	version string = "0.1"
	requestTimeout time.Duration = 7000 * time.Millisecond

)

func New(clientId,  string) BrightpearlAPI{
	ApiInstance := BrightpearlAPI{}
	return ApiInstance
}

func (b BrightpearlAPI) authroizationUrl() string{

	return ""
}


func (b BrightpearlAPI) fetchOauthToken() string{
	return ""
}

func (b BrightpearlAPI) refreshToken() string{
	return ""
}




