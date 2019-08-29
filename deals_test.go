package hubspot

import (
	"os"
	"strconv"
	"testing"
)

func TestDeals(t *testing.T) {
	tvid, err := strconv.Atoi(os.Getenv("HUBSPOT_TEST_VID"))
	if err != nil {
		t.Fatal(err)
	}

	data := DealsRequest{
		Associations: Associations{
			AssociatedVids: []int{tvid},
		},
		Properties: []Property{
			Property{
				Name:  "dealname",
				Value: "Test Deal Activity",
			},
			Property{
				Name:  "brigade_country",
				Value: "Honduras",
			},
			Property{
				Name:  "brigade_name",
				Value: "Test Brigade",
			},
			Property{
				Name:  "brigade_leader_name",
				Value: "Test Brigade Leader",
			},
		},
	}
	c := NewClient(NewClientConfig())
	r, err := c.Deals().Create(data)
	if err != nil {
		t.Error(err)
	}
	if r.ErrorResponse.Status == "error" {
		t.Error(r.ErrorResponse.Message)
	}

	if r.DealID != 0 {
		r, err = c.Deals().Get(r.DealID)
		if err != nil {
			t.Error(err)
		}
		if r.ErrorResponse.Status == "error" {
			t.Error(r.ErrorResponse.Message)
		}
	}

	if r.DealID != 0 {
		err = c.Deals().Delete(r.DealID)
		if err != nil {
			t.Error(err)
		}
	}

	//t.Logf("%+v", r)
}
