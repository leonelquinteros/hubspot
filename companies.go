package hubspot

import "strconv"

// Companies client
type Companies struct {
	Client
}

// Companies constructor (from Client)
func (c Client) Companies() Companies {
	return Companies{
		Client: c,
	}
}

// CompaniesRequest object
type CompaniesRequest struct {
	Properties []Property `json:"properties"`
}

// CompaniesResponse object
type CompaniesResponse struct {
	ErrorResponse
	CompanyID  int               `json:"companyId"`
	PortalID   int               `json:"portalId"`
	Properties CompanyProperties `json:"properties"`
	IsDeleted  bool              `json:"isDeleted"`
}

// CompanyProperties response object
type CompanyProperties struct {
	CreateDate  ResponseProperty `json:"createdate"`
	Name        ResponseProperty `json:"name"`
	Description ResponseProperty `json:"description"`
}

// Get Company
func (c Companies) Get(companyID int) (CompaniesResponse, error) {
	r := CompaniesResponse{}
	err := c.Client.Request("GET", "/companies/v2/companies/"+strconv.Itoa(companyID), nil, &r)
	return r, err
}

// Create new Company
func (c Companies) Create(data CompaniesRequest) (CompaniesResponse, error) {
	r := CompaniesResponse{}
	err := c.Client.Request("POST", "/companies/v2/companies/", data, &r)
	return r, err
}

// Update Deal
func (c Companies) Update(companyID int, data CompaniesRequest) (CompaniesResponse, error) {
	r := CompaniesResponse{}
	err := c.Client.Request("PUT", "/companies/v2/companies/"+strconv.Itoa(companyID), data, &r)
	return r, err
}

// Delete Deal
func (c Companies) Delete(companyID int) error {
	err := c.Client.Request("DELETE", "/companies/v2/companies/"+strconv.Itoa(companyID), nil, nil)
	return err
}
