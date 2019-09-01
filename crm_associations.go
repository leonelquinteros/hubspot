package hubspot

const (
	CRMAssociationContactToCompany            = 1
	CRMAssociationCompanyToContact            = 2
	CRMAssociationDealToContact               = 3
	CRMAssociationContactToDeal               = 4
	CRMAssociationDealToCompany               = 5
	CRMAssociationCompanyToDeal               = 6
	CRMAssociationCompanyToEngagement         = 7
	CRMAssociationEngagementToCompany         = 8
	CRMAssociationContactToEngagement         = 9
	CRMAssociationEngagementToContact         = 10
	CRMAssociationDealToEngagement            = 11
	CRMAssociationEngagementToDeal            = 12
	CRMAssociationParentCompanyToChildCompany = 13
	CRMAssociationChildCompanyToParentCompany = 14
	CRMAssociationContactToTicket             = 15
	CRMAssociationTicketToContact             = 16
	CRMAssociationTicketToEngagement          = 17
	CRMAssociationEngagementToTicket          = 18
	CRMAssociationDealToLineItem              = 19
	CRMAssociationLineItemToDeal              = 20
	CRMAssociationCompanyToTicket             = 25
	CRMAssociationTicketToCompany             = 26
	CRMAssociationDealToTicket                = 27
	CRMAssociationTicketToDeal                = 28
	CRMAssociationAdvisorToCompany            = 33
	CRMAssociationCompanyToAdvisor            = 34
	CRMAssociationBoardMemberToCompany        = 35
	CRMAssociationCompanyToBoardMember        = 36
	CRMAssociationContractorToCompany         = 37
	CRMAssociationCompanyToContractor         = 38
	CRMAssociationManagerToCompany            = 39
	CRMAssociationCompanyToManager            = 40
	CRMAssociationBusinessOwnerToCompany      = 41
	CRMAssociationCompanyToBusinessOwner      = 42
	CRMAssociationPartnerToCompany            = 43
	CRMAssociationCompanyToPartner            = 44
	CRMAssociationResellerToCompany           = 45
	CRMAssociationCompanyToReseller           = 46
)

// CRMAssociations API client
type CRMAssociations struct {
	Client
}

// CRMAssociations constructor (from Client)
func (c Client) CRMAssociations() CRMAssociations {
	return CRMAssociations{
		Client: c,
	}
}

// CRMAssociationsRequest object
type CRMAssociationsRequest struct {
	FromObjectID int    `json:"fromObjectId"`
	ToObjectID   int    `json:"toObjectId"`
	Category     string `json:"category"`
	DefinitionID int    `json:"definitionId"`
}

// Create new CRM Association
func (ca CRMAssociations) Create(data CRMAssociationsRequest) error {
	// Validate Category
	if data.Category == "" {
		data.Category = "HUBSPOT_DEFINED"
	}

	return ca.Client.Request("PUT", "/crm-associations/v1/associations", data, nil)
}

// Delete CRM Association
func (ca CRMAssociations) Delete(data CRMAssociationsRequest) error {
	// Validate Category
	if data.Category == "" {
		data.Category = "HUBSPOT_DEFINED"
	}

	return ca.Client.Request("PUT", "/crm-associations/v1/associations/delete", data, nil)
}
