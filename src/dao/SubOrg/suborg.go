package suborg

type SubOrg struct {
	SubOrgID int
	SubOrgCategoryCodeID   int
	SubOrgTypeCode         string
	SubOrgName             string
	SubOrgDescription      string
	SubOrgPrimaryContact   string
	SubOrgAddressID        int
}

type SubOrgToBuilding struct {
	SubOrganizationID int
	BuildingID int
}