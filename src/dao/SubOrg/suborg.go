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

alter table sub_organization   add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table sub_organization    add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table sub_organization_to_address   add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table sub_organization_to_address     add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table sub_organization_to_building   add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table sub_organization_to_building     add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();


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
	Created  string
	Updated  string
}

type SubOrgToBuilding struct {
	Id int
	SubOrganizationID int
	BuildingID int
	Created  string
	Updated  string
}