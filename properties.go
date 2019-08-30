package hubspot

// Property object is a general property definition.
// To handle inconsistencies in the API, it will take Name or Property to indicate the Property Name and a Value.
type Property struct {
	Name     string      `json:"name,omitempty"`
	Property string      `json:"property,omitempty"`
	Value    interface{} `json:"value"`
}

// ResponseProperty object is the format on which response objects format their property information.
type ResponseProperty struct {
	Name      string             `json:"name"`
	Value     string             `json:"value"`
	Timestamp int64              `json:"timestamp"`
	Source    string             `json:"source"`
	SourceID  string             `json:"sourceId"`
	SourceVid []int              `json:"sourceVid"`
	Versions  []ResponseProperty `json:"versions"`
}
