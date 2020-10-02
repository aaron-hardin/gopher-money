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
	Rates      map[string]float64
}

func NewApiClient(apiKey string) *ApiClient {
	c := new(ApiClient)
	c.ApiKey = apiKey
	return c
}

type ApiClient struct {
	ApiKey string
}

func (c *ApiClient) GetRates() (map[string]float64, error) {
	url := "https://openexchangerates.org/api/latest.json?app_id=" + c.ApiKey
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("rates: error connecting to openexchangerates.org: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("rates: error reading response from openexchangerates.org: %w", err)
	}

	var parsedResponse ratesResponse
	err = json.Unmarshal(body, &parsedResponse)
	if err != nil {
		return nil, fmt.Errorf("rates: error converting response from openexchangerates.org to rates: %w, response: %v", err, string(body))
	}

	return parsedResponse.Rates, nil
}
