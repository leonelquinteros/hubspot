package hubspot

// Association object
type Association struct {
	FromObjectID    int    `json:"fromObjectId"`
	AssociationType string `json:"associationType"`
	ToObjectID      int    `json:"toObjectId"`
	Timestamp       int64  `json:"timestamp"`
}

// Associations object
type Associations struct {
	AssociatedCompanyIds []int `json:"associatedCompanyIds,omitempty"`
	AssociatedVids       []int `json:"associatedVids,omitempty"`
	AssociatedDealIds    []int `json:"associatedDealIds,omitempty"`
}

// AssociationCreateFailure object
type AssociationCreateFailure struct {
	Association Association `json:"association"`
	FailReason  string      `json:"failReason"`
	Message     string      `json:"message"`
}
