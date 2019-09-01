// Package hubspot is a SDK client library for HubSpot's API
//
// TODO: Write docs here...
package hubspot

import "os"

var (
	apiHost    = "https://api.hubapi.com"
	apiKey     string
	oauthToken string
)

func init() {
	// Set environment variables configuration.
	if os.Getenv("HUBSPOT_API_HOST") != "" {
		apiHost = os.Getenv("HUBSPOT_API_HOST")
	}
	if os.Getenv("HUBSPOT_API_KEY") != "" {
		apiKey = os.Getenv("HUBSPOT_API_KEY")
	}
	if os.Getenv("HUBSPOT_OAUTH_TOKEN") != "" {
		oauthToken = os.Getenv("HUBSPOT_OAUTH_TOKEN")
	}
}

// ErrorResponse object
type ErrorResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
