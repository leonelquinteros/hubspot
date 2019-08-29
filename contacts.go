package hubspot

import "strconv"

// Contacts client
type Contacts struct {
	Client
}

// Contacts constructor (from Client)
func (c Client) Contacts() Contacts {
	return Contacts{
		Client: c,
	}
}

// ContactsRequest object
type ContactsRequest struct {
	Properties []Property `json:"properties"`
}

// ContactsResponse object
type ContactsResponse struct {
	ErrorResponse
	PortalID          int                        `json:"portal-id"`
	Vid               int                        `json:"vid"`
	CanonicalVid      int                        `json:"canonical-vid"`
	MergeVids         []int                      `json:"merge-vids"`
	IsContact         bool                       `json:"is-contact"`
	ProfileToken      string                     `json:"profile-token"`
	ProfileURL        string                     `json:"profile-url"`
	IdentityProfiles  []IdentityProfile          `json:"identity-profiles"`
	Properties        map[string]ContactProperty `json:"properties"`
	FormSubmissions   []interface{}              `json:"form-submissions"`
	AssociatedCompany AssociatedCompany          `json:"associated-company"`
}

// AssociatedCompany object
type AssociatedCompany struct {
	CompanyID  int        `json:"company-id"`
	PortalID   int        `json:"portal-id"`
	Properties []Property `json:"properties"`
}

// CreateOrUpdateContactResponse object
type CreateOrUpdateContactResponse struct {
	ErrorResponse
	Vid   int  `json:"vid"`
	IsNew bool `json:"isNew"`
}

// DeleteContactResponse object
type DeleteContactResponse struct {
	ErrorResponse
	Vid     int    `json:"vid"`
	Deleted bool   `json:"deleted"`
	Reason  string `json:"reason"`
}

// IdentityProfile response object
type IdentityProfile struct {
	Identities []struct {
		Timestamp int64  `json:"timestamp"`
		Type      string `json:"type"`
		Value     string `json:"value"`
	} `json:"identities"`
	Vid int `json:"vid"`
}

// ContactProperty response object
type ContactProperty struct {
	Value    string `json:"value"`
	Versions []struct {
		Value       string      `json:"value"`
		Timestamp   int64       `json:"timestamp"`
		SourceType  string      `json:"source-type"`
		SourceID    interface{} `json:"source-id"`
		SourceLabel interface{} `json:"source-label"`
		Selected    bool        `json:"selected"`
	} `json:"versions"`
}

// Get Contact
func (c Contacts) Get(contactID int) (ContactsResponse, error) {
	r := ContactsResponse{}
	err := c.Client.Request("GET", "/contacts/v1/contact/vid/"+strconv.Itoa(contactID)+"/profile", nil, &r)
	return r, err
}

// GetByEmail a Contact
func (c Contacts) GetByEmail(email string) (ContactsResponse, error) {
	r := ContactsResponse{}
	err := c.Client.Request("GET", "/contacts/v1/contact/email/"+email+"/profile", nil, &r)
	return r, err
}

// Create new Contact
func (c Contacts) Create(data ContactsRequest) (ContactsResponse, error) {
	r := ContactsResponse{}
	err := c.Client.Request("POST", "/contacts/v1/contact", data, &r)
	return r, err
}

// Update Contact
func (c Contacts) Update(contactID int, data ContactsRequest) error {
	return c.Client.Request("POST", "/contacts/v1/contact/vid/"+strconv.Itoa(contactID)+"/profile", data, nil)
}

// UpdateByEmail a Contact
func (c Contacts) UpdateByEmail(email string, data ContactsRequest) error {
	return c.Client.Request("POST", "/contacts/v1/contact/email/"+email+"/profile", data, nil)
}

// CreateOrUpdate a Contact
func (c Contacts) CreateOrUpdate(email string, data ContactsRequest) (CreateOrUpdateContactResponse, error) {
	r := CreateOrUpdateContactResponse{}
	err := c.Client.Request("POST", "/contacts/v1/contact/createOrUpdate/email/"+email, data, &r)
	return r, err
}

// Delete Deal
func (c Contacts) Delete(contactID int) error {
	return c.Client.Request("DELETE", "/contacts/v1/contact/vid/"+strconv.Itoa(contactID), nil, nil)
}
