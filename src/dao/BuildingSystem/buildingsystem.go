package BuildingSystem

/**
CREATE TABLE building_system (
	building_system_id                   integer PRIMARY KEY NOT NULL,
	building_system_category_code          integer  NOT NULL,
	building_system_type                  varchar(50) NOT NULL,
	building_system_name                  varchar(50) NOT NULL,
	building_system_description           varchar(256) NOT NULL,
    building_system_is_active              BOOLEAN NOT NULL
)
*/
type BuildingSystem struct {
	BuildingSystemId          int
	BuildingSystemName        string
	BuildingSystemType        string
	BuildingSystemDescription string
	isActive                  string
}
