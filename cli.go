package go_waka_api

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	API_ENDPOINT string = "https://wakatime.com/api/v1/"
)

// Create new instance
func New(token string) *WakaClient {
	return &WakaClient{
		client: &http.Client{
			Timeout: time.Second * 15,
			Transport: &http.Transport{
				IdleConnTimeout: time.Minute * 2,
				MaxIdleConns:    10,
			},
		},
		token: base64.StdEncoding.EncodeToString([]byte(token)),
	}
}

func sendRequest(waka *WakaClient, request *http.Request) ([]byte, error) {
	request.Header.Add("Authorization", "Basic "+waka.token)
	res, err := waka.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
		break
	default:
		return nil, fmt.Errorf("Status code: %d.\nMessage: %s", res.StatusCode, string(data))
	}

	return data, nil
}

func (w *WakaClient) get(link string, query map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, API_ENDPOINT+link, nil)
	if err != nil {
		return nil, err
	}
	if query != nil {
		q := req.URL.Query()
		for key, value := range query {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}
	return sendRequest(w, req)
}

func (w *WakaClient) post(link string, json []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, API_ENDPOINT+link, bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}
	return sendRequest(w, req)
}
