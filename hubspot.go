/*
Package hubspot is a SDK client library for HubSpot's API

Example usage:

	package main

	import "github.com/leonelquinteros/hubspot"

	func main() {
		// Create client object with config from environment variables (HUBSPOT_API_HOST, HUBSPOT_API_KEY, HUBSPOT_OAUTH_TOKEN)
		c := hubspot.NewClient(hubspot.NewClientConfig())

		// Create new contact
		data := hubspot.ContactsRequest{
			hubspot.Properties: []hubspot.Property{
				hubspot.Property{
					Property: "email",
					Value:    "contact@example.com",
				},
				hubspot.Property{
					Property: "firstname",
					Value:    "Contact",
				},
				hubspot.Property{
					Property: "lastname",
					Value:    "Example",
				},
			},
		}
		r, err := c.Contacts().Create(data)
		if err != nil {
			log.Fatal(err)
		}

		// Get contact by email
		contact, err = c.Contacts().GetByEmail("contact@example.com")
		if err != nil {
			log.Fatal(err)
		}

		// Print contact object
		fmt.Printf("%+v", contact)
	}


Configuration

Package can be configured using environment variables for default initialization.
When calling

	hubspot.NewClientConfig()

the configuration object will be populated with information from environment variables as follow:

	HUBSPOT_API_HOST

The HubSpot's API endpoint host (including schema and host), defaults to "https://api.hubapi.com" when omitted.

	HUBSPOT_API_KEY

HubSpot's API secret key used to authenticate further requests. It will be preferred if both auth methods (API key and OAuth token) are present.

	HUBSPOT_OAUTH_TOKEN

If HUBSPOT_API_KEY isn't present, further requests will be authenticated using this token.

Manual/eventual configuration can be set by creating a configuration object directly and pass it to the NewClient() method:

	hubspot.NewClient(hubspot.ClientConfig{
		APIHost: "https://api.hubspot.local"
		APIKey: "vroho487hfo48hfo48y3bai38gi2"
	})


*/
package hubspot

// ErrorResponse object
type ErrorResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
