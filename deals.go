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
	PortalID                  int                      `json:"portalId"`
	DealID                    int                      `json:"dealId"`
	IsDeleted                 bool                     `json:"isDeleted"`
	Associations              Associations             `json:"associations"`
	AssociationCreateFailures AssociationCreateFailure `json:"associationCreateFailures"`
	Properties                DealProperties           `json:"properties"`
}

// DealProperties object
type DealProperties struct {
	Amount         DealProperty `json:"amount"`
	Dealstage      DealProperty `json:"dealstage"`
	Pipeline       DealProperty `json:"pipeline"`
	Closedate      DealProperty `json:"closedate"`
	Createdate     DealProperty `json:"createdate"`
	HubspotOwnerID DealProperty `json:"hubspot_owner_id"`
	HsCreatedate   DealProperty `json:"hs_createdate"`
	Dealtype       DealProperty `json:"dealtype"`
	Dealname       DealProperty `json:"dealname"`
}

// DealProperty object
type DealProperty struct {
	Name      string         `json:"name"`
	Value     string         `json:"value"`
	Timestamp int64          `json:"timestamp"`
	Source    string         `json:"source"`
	SourceID  string         `json:"sourceId"`
	SourceVid []int          `json:"sourceVid"`
	Versions  []DealProperty `json:"versions"`
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
