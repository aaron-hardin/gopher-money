package rates

import (
	"os"
	"testing"
)

func TestApiClient_GetRates_Integration(t *testing.T) {
	// Marked as integration test because it requires network connection
	// to Open Exchange Rates API
	if testing.Short() {
		t.Skip("Skipping test because it requires network")
	}

	// Fail test if not set, with descriptive err
	apiKey := os.Getenv("GOPHER_MONEY_API_KEY")
	if apiKey == "" {
		t.Fatal("Could not get environment variable required for test: GOPHER_MONEY_API_KEY")
	}

	client := NewApiClient(apiKey)
	rates, err := client.GetRates()
	if err != nil {
		t.Fatal(err)
	}

	if len(rates) < 1 {
		t.Fatal("No rates returned from API")
	}
}
