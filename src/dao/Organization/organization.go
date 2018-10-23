package Organization

/*
CREATE TABLE organization (
    organization_id                   integer PRIMARY KEY NOT NULL,
	organization_category_code_id     integer  NOT NULL,
	organization_type                  varchar(50) NOT NULL,
	organization_name                  varchar(50) NOT NULL,
	organization_description           varchar(256) NOT NULL,
    organization_address_id              integer NOT NULL,
	organization_primary_contact_name         varchar(50),
    organization_primary_contact_email         varchar(50)
)


ALTER  TABLE organization DROP COLUMN organization_category_code_id;
alter table organization add column organization_category_code  varchar(50) NOT NULL;

create table organization_to_address
    (
	  id                   integer PRIMARY KEY NOT NULL,
      organization_id          integer REFERENCES organization NOT NULL,
      address_id   integer  REFERENCES address NOT NULL
	)

create table organization_to_suborg
    (
	  id                   integer PRIMARY KEY NOT NULL,
      organization_id          integer REFERENCES organization NOT NULL,
      sub_organization_id   integer  REFERENCES sub_organization NOT NULL
	)

create table organization_to_building
    (
	  id                   integer PRIMARY KEY NOT NULL,
      organization_id          integer REFERENCES organization NOT NULL,
      building_id   integer  REFERENCES building NOT NULL
	)
 */

type Organization struct {
	OrganizationID int
	OrganizationCategoryCodeID   int
	OrganizationType             string
    OrganizationName             string
	OrganizationDescription      string
	OrganizationPrimaryContactName   string
	OrganizationPrimaryEmail         string
	OrganizationAddressID             int
}

type OrgToSubOrg struct {
	OrganizationID int
	SubOrganizqationID int
}

type OrgToBuilding struct {
	OrganizationID int
	BuildingID int
}