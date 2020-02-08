package brightpearl

import (
"fmt"
"github.com/shriharishastry/brightpearl/connector"
"net/http"
"net/url"
"time"
)

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
	requestTimeout time.Duration = 7000 * time.Millisecond
	AuthTokenURL string = "https://oauth.brightpearl.com/authorize"
	FetchTokenURL string = "https://oauth.brightpearl.com/token"
	RefreshTokenURL string = "https://oauth.brightpearl.com/token"

)

func New(ClientId, ClientSecret, Domain, AppRef, DeveloperRef, Account string) BrightpearlAPI{
	ApiInstance := BrightpearlAPI{
		ClientId:ClientId,
		ClientSecret:ClientSecret,
		Domain:Domain,
		AppRef:AppRef,
		DeveloperRef:DeveloperRef,
		Account:Account,
	}
	return ApiInstance
}

func (b BrightpearlAPI) GetAuthURL() string{
	queryParams := url.Values{}
	queryParams.Add("response_type", "code")
	queryParams.Add("client_id", b.ClientId)
	queryParams.Add("redirect_url", "")
	queryParams.Add("state", "random")
	return fmt.Sprintf("%s/%s?%s", AuthTokenURL, b.Account, queryParams.Encode())
}


func (b BrightpearlAPI) FetchOauthToken() string{
	conn := connector.HttpClient{}
	resp, err := conn.SendRequest(http.MethodGet, fmt.Sprintf("%s", FetchTokenURL), nil, nil)
	fmt.Println(resp, err)
	return ""
}

func (b BrightpearlAPI) RefreshToken() string{
	conn := connector.HttpClient{}
	resp, err := conn.SendRequest(http.MethodGet, fmt.Sprintf("%s", RefreshTokenURL), nil, nil)
	fmt.Println(resp, err)
	return ""
}



