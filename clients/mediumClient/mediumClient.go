package mediumClient

import (
	"encoding/xml"
	"errors"
	client "github.com/bozd4g/go-http-client"
	"log"
	"os"
)

type mediumClient struct {
	client client.IHttpClient
}

type IMediumClient interface {
	GetRss() (*RssDto, error)
}

func New() IMediumClient {
	apiBaseUrl := os.Getenv("API_BASE_URL")

	httpClient := client.New(apiBaseUrl)
	return &mediumClient{client: httpClient}
}

func (m *mediumClient) GetRss() (*RssDto, error) {
	apiGetMethod := os.Getenv("API_GET_METHOD")
	if apiGetMethod == "" {
		return nil, errors.New("API_GET_METHOD cannot be empty")
	}

	request, err := m.client.Get(apiGetMethod)
	if err != nil {
		log.Println("Error: " + err.Error())
	}

	var rss *RssDto
	response := m.client.Do(request)
	if response.IsSuccess {
		err = xml.Unmarshal([]byte(response.Data), &rss)
	} else {
		err = errors.New(response.Message)
		log.Println("Error: " + response.Message)
	}

	return rss, err
}
