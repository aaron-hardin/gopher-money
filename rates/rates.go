package rates

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ratesResponse struct {
	Disclaimer string
	License    string
	Timestamp  uint32 // Valid until 2106
	Base       string
	Rates      map[string]float32
}

func NewApiClient(apiKey string) *ApiClient {
	c := new(ApiClient)
	c.ApiKey = apiKey
	return c
}

type ApiClient struct {
	ApiKey string
}

func (c *ApiClient) GetRates() map[string]float32 {
	url := "https://openexchangerates.org/api/latest.json?app_id=" + c.ApiKey
	resp, err := http.Get(url)
	if err != nil {
		// TODO: handle error
		fmt.Println("error:", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO: handle error
		fmt.Println("error:", err)
	}

	var parsedResponse ratesResponse
	err = json.Unmarshal(body, &parsedResponse)
	if err != nil {
		// TODO: handle error
		fmt.Println("error:", err)
	}

	return parsedResponse.Rates
}
