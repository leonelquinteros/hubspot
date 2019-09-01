package hubspot

import "testing"

func TestCRMAssociations(t *testing.T) {
	// Init client
	c := NewClient(NewClientConfig())

	// Create company
	companyData := CompaniesRequest{
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
	companyResp, err := c.Companies().Create(companyData)
	if err != nil {
		t.Error(err)
	}
	if companyResp.ErrorResponse.Status == "error" {
		t.Error(companyResp.ErrorResponse.Message)
	}

	// Create contact
	contactData := ContactsRequest{
		Properties: []Property{
			Property{
				Property: "email",
				Value:    "crm_associations@example.com",
			},
			Property{
				Property: "firstname",
				Value:    "CRM",
			},
			Property{
				Property: "lastname",
				Value:    "Associations",
			},
		},
	}
	contactResp, err := c.Contacts().Create(contactData)
	if err != nil {
		t.Fatal(err)
	}
	if contactResp.ErrorResponse.Status == "error" {
		t.Fatal(contactResp.ErrorResponse.Message)
	}

	// Create CRM association contact -> company
	assocData := CRMAssociationsRequest{
		FromObjectID: contactResp.Vid,
		ToObjectID:   companyResp.CompanyID,
		DefinitionID: CRMAssociationContactToCompany,
	}
	err = c.CRMAssociations().Create(assocData)
	if err != nil {
		t.Fatal(err)
	}

	// Delete association
	err = c.CRMAssociations().Delete(assocData)
	if err != nil {
		t.Fatal(err)
	}

	// Delete contact
	err = c.Contacts().Delete(contactResp.Vid)
	if err != nil {
		t.Fatal(err)
	}

	// Delete Company
	err = c.Companies().Delete(companyResp.CompanyID)
	if err != nil {
		t.Error(err)
	}
}
