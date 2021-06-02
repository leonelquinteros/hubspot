package hubspot

import "strconv"

// Deals client
type Deals struct {
	Client
}

// Deals constructor (from Client)
func (c Client) Deals() Deals {
	return Deals{
		Client: c,
	}
}

// DealsRequest object
type DealsRequest struct {
	Associations Associations `json:"associations,omitempty"`
	Properties   []Property   `json:"properties"`
}

// DealsResponse object
type DealsResponse struct {
	ErrorResponse
	PortalID                  int                        `json:"portalId"`
	DealID                    int                        `json:"dealId"`
	IsDeleted                 bool                       `json:"isDeleted"`
	Associations              Associations               `json:"associations"`
	AssociationCreateFailures []AssociationCreateFailure `json:"associationCreateFailures"`
	Properties                DealProperties             `json:"properties"`
}

// DealProperties object
type DealProperties struct {
	Amount         ResponseProperty `json:"amount"`
	Dealstage      ResponseProperty `json:"dealstage"`
	Pipeline       ResponseProperty `json:"pipeline"`
	Closedate      ResponseProperty `json:"closedate"`
	Createdate     ResponseProperty `json:"createdate"`
	HubspotOwnerID ResponseProperty `json:"hubspot_owner_id"`
	HsCreatedate   ResponseProperty `json:"hs_createdate"`
	Dealtype       ResponseProperty `json:"dealtype"`
	Dealname       ResponseProperty `json:"dealname"`
	BuilderToken   ResponseProperty `json:"builder_token"`
}

// Get Deal
func (d Deals) Get(dealID int) (DealsResponse, error) {
	r := DealsResponse{}
	err := d.Client.Request("GET", "/deals/v1/deal/"+strconv.Itoa(dealID), nil, &r)
	return r, err
}

// Create new Deal
func (d Deals) Create(data DealsRequest) (DealsResponse, error) {
	r := DealsResponse{}
	err := d.Client.Request("POST", "/deals/v1/deal/", data, &r)
	return r, err
}

// Update Deal
func (d Deals) Update(dealID int, data DealsRequest) (DealsResponse, error) {
	r := DealsResponse{}
	err := d.Client.Request("PUT", "/deals/v1/deal/"+strconv.Itoa(dealID), data, &r)
	return r, err
}

// Delete Deal
func (d Deals) Delete(dealID int) error {
	err := d.Client.Request("DELETE", "/deals/v1/deal/"+strconv.Itoa(dealID), nil, nil)
	return err
}
