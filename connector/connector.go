package connector


import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type HttpClient struct {
	Client *http.Client
}

type HTTPResponse struct {
	Body     []byte
	Response *http.Response
}

func (c *HttpClient) SendRequest(method string, rURL string, params url.Values, headers http.Header) (HTTPResponse, error){
	var (
		resp = HTTPResponse{}
		postParam io.Reader
		err error
	)

	if params == nil{
		params = url.Values{}
	}

	if method == http.MethodPost || method == http.MethodPut {
		postParam = strings.NewReader(params.Encode())
	}

	request, err:= http.NewRequest(method, rURL, postParam)
	if headers != nil {
		request.Header = headers
	}

	if method == http.MethodGet{
		request.URL.RawQuery = params.Encode()
	}

	if err != nil{
		log.Fatal(err)
		return resp, err
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	resp.Response = response
	resp.Body = body
	return resp, nil

}
//
//func (c *HttpClient) SendAndProcessRequest(method string, rURL string, params url.Values, headers http.Header, dataObj interface{}) error{
//		resp, err := c.SendRequest(method, rURL, params, headers)
//
//
//}
