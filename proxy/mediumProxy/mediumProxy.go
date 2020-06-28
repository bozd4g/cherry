package mediumProxy

import (
	"encoding/json"
	"errors"
	"github.com/bozd4g/cherry/models"
	client "github.com/bozd4g/go-http-client"
	"log"
	"os"
)

type mediumProxy struct {
	client client.IHttpClient
}

type IMediumProxy interface {
	GetRss() (*models.Rss, error)
}

func New() IMediumProxy {
	apiBaseUrl := os.Getenv("API_BASE_URL")

	httpClient := client.New(apiBaseUrl)
	return &mediumProxy{client: httpClient}
}

func (m *mediumProxy) GetRss() (*models.Rss, error) {
	apiGetMethod := os.Getenv("API_GET_METHOD")

	request, err := m.client.Get(apiGetMethod)
	if err != nil {
		log.Println("Error: " + err.Error())
	}

	var rss *models.Rss
	response := m.client.Do(request)
	if response.IsSuccess {
		err = json.Unmarshal([]byte(response.Data), &rss)
	} else {
		err = errors.New(response.Message)
		log.Println("Error: " + response.Message)
	}

	return rss, err
}
