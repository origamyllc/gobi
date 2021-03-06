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

alter table organization   add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table organization    add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table organization_to_address   add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table organization_to_address    add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table organization_to_building  add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table organization_to_building    add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table organization_to_suborg   add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table organization_to_suborg    add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();



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
	Created  string
	Updated  string
}

type OrgToSubOrg struct {
	Id int
	OrganizationID int
	SubOrganizqationID int
	Created  string
	Updated  string
}

type OrgToBuilding struct {
	Id int
	OrganizationID int
	BuildingID int
	Created  string
	Updated  string
}