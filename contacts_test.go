package hubspot

import (
	"testing"

	"github.com/leonelquinteros/gorand"
)

func TestContacts(t *testing.T) {
	tEmailUser, err := gorand.GetAlphaNumString(8)
	if err != nil {
		t.Fatal(err)
	}
	tEmailHost, err := gorand.GetAlphaNumString(6)
	if err != nil {
		t.Fatal(err)
	}
	tEmail := tEmailUser + "@" + tEmailHost + ".com"

	data := ContactsRequest{
		Properties: []Property{
			Property{
				Property: "email",
				Value:    tEmail,
			},
			Property{
				Property: "firstname",
				Value:    "Test",
			},
			Property{
				Property: "lastname",
				Value:    "Contact",
			},
		},
	}

	// Create
	c := NewClient(NewClientConfig())
	r, err := c.Contacts().Create(data)
	if err != nil {
		t.Fatal(err)
	}
	if r.ErrorResponse.Status == "error" {
		t.Fatal(r.ErrorResponse.Message)
	}

	// Update
	data.Properties[1].Value = data.Properties[1].Value.(string) + "1"
	if r.Vid != 0 {
		err = c.Contacts().Update(r.Vid, data)
		if err != nil {
			t.Fatal(err)
		}
		if r.ErrorResponse.Status == "error" {
			t.Fatal(r.ErrorResponse.Message)
		}
	}

	// CreateOrUpdate
	data.Properties[2].Value = data.Properties[2].Value.(string) + "2"
	if r.Vid != 0 {
		ru, err := c.Contacts().CreateOrUpdate(data.Properties[0].Value.(string), data)
		if err != nil {
			t.Fatal(err)
		}
		if ru.ErrorResponse.Status == "error" {
			t.Fatal(r.ErrorResponse.Message)
		}
	}

	// Get contact
	contact, err := c.Contacts().Get(r.Vid)
	if err != nil {
		t.Fatal(err)
	}
	if r.ErrorResponse.Status == "error" {
		t.Fatal(r.ErrorResponse.Message)
	}

	// Get by email
	contact, err = c.Contacts().GetByEmail(data.Properties[0].Value.(string))
	if err != nil {
		t.Fatal(err)
	}
	if r.ErrorResponse.Status == "error" {
		t.Fatal(r.ErrorResponse.Message)
	}

	// Delete
	err = c.Contacts().Delete(contact.Vid)
	if err != nil {
		t.Fatal(err)
	}
	if r.ErrorResponse.Status == "error" {
		t.Fatal(r.ErrorResponse.Message)
	}

	//t.Logf("%+v", r)
	//t.Logf("%+v", contact)
}
