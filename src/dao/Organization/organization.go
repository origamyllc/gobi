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
alter table organization
      add constraint organization_to_sub_org
      foreign key (organization_id)
      references sub_organization (sub_organization_id);

alter table organization
      add constraint organization_to_building
      foreign key (organization_id)
      references building (building_id);
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