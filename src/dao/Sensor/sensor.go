package Sensor

/**
CREATE TABLE sensor (
    sensor_id                   integer PRIMARY KEY NOT NULL,
	sensor_type                  varchar(50) NOT NULL,
	sensor_name                  varchar(50) NOT NULL,
	sensor_make                  varchar(50) NOT NULL,
    sensor_model                  varchar(50) NOT NULL,
    sensor_ip_address             varchar(50) NOT NULL,
    sensor_mac_addess              varchar(50) NOT NULL,
	sensor_description              varchar(256) NOT NULL,
    sensor_location              varchar(50) NOT NULL,
	sensor_data_sheet_location       varchar(512) NOT NULL,
	sensor_average_hourly_energy_consumption     FLOAT(10),
	sensor_average_weekly_energy_consumption     FLOAT(10),
	sensor_average_monthly_energy_consumption    FLOAT(10),
	sensor_average_yearly_energy_consumption     FLOAT(10)
)



create table sensor_to_protocol
    (
	  id                   integer PRIMARY KEY NOT NULL,
      sensor_id          integer REFERENCES sensor NOT NULL,
      protocol_id       integer  REFERENCES protocols NOT NULL
	);

create table sensor_to_building_system
    (
	  id                   integer PRIMARY KEY NOT NULL,
      sensor_id          integer REFERENCES sensor NOT NULL,
      building_system_id       integer  REFERENCES building_system  NOT NULL
	);

create table sensor_to_observable
    (
	  id                   integer PRIMARY KEY NOT NULL,
      sensor_id          integer REFERENCES sensor NOT NULL,
      observable_id       integer  REFERENCES observable NOT NULL
	);

alter table sensor   add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table sensor    add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table sensor_to_protocol   add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table sensor_to_protocol     add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table sensor_to_building_system   add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table sensor_to_building_system     add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table sensor_to_observable add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table sensor_to_observable  add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
 */

type Sensor struct {
	SensorID   string
	SensorName string
	SensorMake string
	SensorModel string
	SensorType string
	SensorIPAddress string
	SensorMacAddress string
	SensorLocation string
	SensorDescription string
	SensorDatasheetLocation string
	SensorAverageHourlyEnergyConsumption     float64
	SensorAverageWeeklyEnergyConsumption     float64
	SensorAverageMonthlyEnergyConsumption     float64
	SensorAverageYearlyEnergyConsumption     float64
}

type SensorToObservable struct {
	SensorID       int
	ObservableID   int
}

type SensorToProtocol struct {
	SensorID       int
	ProtocolID   int
}

type SensorToBuildingSystem struct {
	SensorID int
	BuildingSystemID    int
}