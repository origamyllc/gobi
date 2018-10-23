package Building

/**

DROP TABLE IF EXISTS building;

CREATE TABLE building (
    building_id                   integer PRIMARY KEY NOT NULL,
	building_type                  varchar(50) NOT NULL,
	building_name                  varchar(50) NOT NULL,
	building_addressid             integer NOT NULL,
	building_year_built             varchar(4) NOT NULL,
	building_description           varchar(256) NOT NULL,
	building_totlal_no_of_people      integer NOT NULL,
	building_total_area             bigint NOT NULL,
	building_total_no_of_floors       integer NOT NULL,
	building_hours_occupied         varchar(50),
	building_average_hourly_energy_consumption     FLOAT(10),
	building_average_weekly_energy_consumption     FLOAT(10),
	building_average_monthly_energy_consumption    FLOAT(10),
	building_average_yearly_energy_consumption     FLOAT(10)
);

CREATE TABLE building_to_floor (
    id                   integer PRIMARY KEY NOT NULL,
    building_id          integer REFERENCES building NOT NULL,
    floor_id             integer  REFERENCES floor NOT NULL
);

CREATE TABLE building_to_building_system (
    id                   integer PRIMARY KEY NOT NULL,
    building_id          integer REFERENCES building NOT NULL,
    building_system_id   integer  REFERENCES building_system NOT NULL
);

create table building_to_address
    (
	  id                   integer PRIMARY KEY NOT NULL,
      building_id   integer  REFERENCES building NOT NULL,
		address_id          integer REFERENCES address NOT NULL
	)

alter table building add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table building add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table building_to_address add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table building_to_address  add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table building_to_building_system  add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table building_to_building_system   add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table building_to_floor  add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table building_to_floor   add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
 */

type Building struct {
	BuildingID                    int
	BuildingType                  string
	BuildingName                  string
	BuildingAddressID             int
	BuildingYearBuilt             string
	BuildingDescription           string
	BuildingTotlalNoOfPeople      int
	BuildingTotalArea             int
	BuildingTotalNoOfFloors       int
	BuildingHoursOccupied         string
	BuildingAverageHourlyEnergyConsumption     float64
	BuildingAverageWeeklyEnergyConsumption     float64
	BuildingAverageMonthlyEnergyConsumption    float64
	BuildingAverageYearlyEnergyConsumption     float64
	Created  string
	Updated  string
}

type BuildingToFloor struct{
	Id   int
	BuildingID int
	FloorID    int
	Created  string
	Updated  string
}

type BuildingToBuildingSystem struct {
	Id   int
	BuildingID int
	BuildingSystemID    int
	Created  string
	Updated  string
}