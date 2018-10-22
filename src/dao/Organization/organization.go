package Organization

type Organization struct {
	OrganizationID int
	OrganizationCategoryCodeID   int
	OrganizationTypeCode         string
    OrganizationName string
	OrganizationDescription      string
	OrganizationPrimaryContact   string
	OrganizationAddressID        int
}

type OrgToCategory struct {
	OrganizationID int
	CategoryCodeID int
}

type OrgToSubOrg struct {
	OrganizationID int
	SubOrganizqationID int
}

type OrgToBuilding struct {
	OrganizationID int
	BuildingID int
}