package hubspot

// Property object is a general property definition.
// To handle inconsistencies in the API, it will take Name or Property to indicate the Property Name and a Value.
type Property struct {
	Name     string      `json:"name,omitempty"`
	Property string      `json:"property,omitempty"`
	Value    interface{} `json:"value"`
}
