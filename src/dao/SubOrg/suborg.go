package suborg
/*
CREATE TABLE sub_organization (
	sub_organization_id                   integer PRIMARY KEY NOT NULL,
	sub_organization_category_code_id     integer  NOT NULL,
	sub_organization_type                  varchar(50) NOT NULL,
	sub_organization_name                  varchar(50) NOT NULL,
	sub_organization_description           varchar(256) NOT NULL,
	sub_organization_address_id              integer NOT NULL,
	sub_organization_primary_contact_name         varchar(50),
	sub_organization_primary_contact_email         varchar(50)
)

create table sub_organization_to_address
    (
	  id                   integer PRIMARY KEY NOT NULL,
      sub_organization_id          integer REFERENCES sub_organization NOT NULL,
      address_id   integer  REFERENCES address NOT NULL
	)
create table sub_organization_to_building
    (
	  id                   integer PRIMARY KEY NOT NULL,
      sub_organization_id          integer REFERENCES sub_organization NOT NULL,
      building_id   integer  REFERENCES building NOT NULL
	)

*/


type SubOrg struct {
	SubOrgID int
	SubOrgCategoryCodeID   int
	SubOrgTypeCode         string
	SubOrgName             string
	SubOrgDescription      string
	SubOrgPrimaryContact   string
	SubOrgAddressID        int
	SubOrganizationPrimaryContactName   string
	SubOrganizationPrimaryEmail         string
}

type SubOrgToBuilding struct {
	SubOrganizationID int
	BuildingID int
}