package hubspot

import "testing"

func TestCompanies(t *testing.T) {
	data := CompaniesRequest{
		Properties: []Property{
			Property{
				Name:  "name",
				Value: "Test Company",
			},
			Property{
				Name:  "description",
				Value: "Company description here...",
			},
		},
	}
	c := NewClient(NewClientConfig())
	r, err := c.Companies().Create(data)
	if err != nil {
		t.Error(err)
	}
	if r.ErrorResponse.Status == "error" {
		t.Error(r.ErrorResponse.Message)
	}

	if r.CompanyID != 0 {
		comp, err := c.Companies().Get(r.CompanyID)
		if err != nil {
			t.Error(err)
		}
		if r.ErrorResponse.Status == "error" {
			t.Error(r.ErrorResponse.Message)
		}

		if comp.CompanyID != 0 {
			data.Properties[1].Value = data.Properties[1].Value.(string) + " updated..."
			r, err := c.Companies().Update(comp.CompanyID, data)
			if err != nil {
				t.Error(err)
			}
			if r.ErrorResponse.Status == "error" {
				t.Error(r.ErrorResponse.Message)
			}
		}
	}

	if r.CompanyID != 0 {
		err = c.Companies().Delete(r.CompanyID)
		if err != nil {
			t.Error(err)
		}
	}

	//t.Logf("%+v", r)
}
