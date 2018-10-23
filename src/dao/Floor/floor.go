package Floor

/*
CREATE TABLE floor (
    floor_id                   integer PRIMARY KEY NOT NULL,
	floor_type                  varchar(50) NOT NULL,
	floor_name                  varchar(50) NOT NULL,
	floor_description           varchar(256) NOT NULL,
	floor_totlal_no_of_people      integer NOT NULL,
	floor_total_area             bigint NOT NULL,
	floor_hours_occupied         varchar(50),
	floors_average_hourly_energy_consumption     FLOAT(10),
	floors_average_weekly_energy_consumption     FLOAT(10),
	floors_average_monthly_energy_consumption    FLOAT(10),
	floors_average_yearly_energy_consumption     FLOAT(10)
)

create table floor_to_building_system
    (
	  id                   integer PRIMARY KEY NOT NULL,
      floor_id          integer REFERENCES floor NOT NULL,
      building_system_id   integer  REFERENCES building_system NOT NULL
	)

create table floor_to_room
    (
	  id                   integer PRIMARY KEY NOT NULL,
      floor_id          integer REFERENCES floor NOT NULL,
      room_id   integer  REFERENCES room NOT NULL
	)

create table floor_to_gateway
    (
	  id                   integer PRIMARY KEY NOT NULL,
      floor_id          integer REFERENCES floor NOT NULL,
      gateway_id   integer  REFERENCES gateway NOT NULL
	)

 */
type Floor struct {
	FloorID                                 int
	FloorType                               int
	FloorName                               int
	FloorDescription                        string
	FloorTotlalNoOfPeople                   int
	FloorTotalArea                          int
	FloorHoursOccupied                      string
	FloorAverageHourlyEnergyConsumption     float64
	FloorAverageWeeklyEnergyConsumption     float64
	FloorAverageMonthlyEnergyConsumption    float64
	FloorAverageYearlyEnergyConsumption     float64
}

type FloorToBuildingSystem struct {
	FloorID   int
	BuildingSystemID    int
}

type FloorToRoom struct {
	FloorID   int
	RoomID    int
}

type FloorToGateway struct {
	FloorID   int
	GatewayID    int
}