package Device


/**
CREATE TABLE device (
    device_id                   integer PRIMARY KEY NOT NULL,
	device_type                  varchar(50) NOT NULL,
	device_name                  varchar(50) NOT NULL,
	device_make                  varchar(50) NOT NULL,
    device_model                  varchar(50) NOT NULL,
    device_ip_address             varchar(50) NOT NULL,
    device_mac_addess              varchar(50) NOT NULL,
	device_description              varchar(256) NOT NULL,
    device_location              varchar(50) NOT NULL,
	device_data_sheet_location       varchar(512) NOT NULL,
	device_average_hourly_energy_consumption     FLOAT(10),
	device_average_weekly_energy_consumption     FLOAT(10),
	device_average_monthly_energy_consumption    FLOAT(10),
	device_average_yearly_energy_consumption     FLOAT(10)
)

create table device_to_sensor
    (
	  id                   integer PRIMARY KEY NOT NULL,
      device_id          integer REFERENCES device NOT NULL,
      sensor_id   integer  REFERENCES sensor NOT NULL
	)

create table device_to_protocol
    (
	  id                   integer PRIMARY KEY NOT NULL,
      device_id          integer REFERENCES device NOT NULL,
      protocol_id   integer  REFERENCES protocol NOT NULL
	)

create table device_to_building_system
    (
	  id                   integer PRIMARY KEY NOT NULL,
      device_id          integer REFERENCES device NOT NULL,
      building_system_id   integer  REFERENCES building_system NOT NULL
	)

alter table device  add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table device   add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table device_to_building_system  add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table device_to_building_system   add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table device_to_protocol  add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table device_to_protocol   add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table device_to_sensor  add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table device_to_sensor   add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();

 */

type Device struct {
	DeviceID   string
    DeviceName string
	DeviceMake string
	DeviceModel string
	DeviceType string
	DeviceIPAddress string
	DeviceMacAddress string
	DeviceLocation string
	DeviceDescription string
	DeviceDatasheetLocation string
	DeviceAverageHourlyEnergyConsumption     float64
	DeviceAverageWeeklyEnergyConsumption     float64
	DeviceAverageMonthlyEnergyConsumption     float64
	DeviceAverageYearlyEnergyConsumption     float64
	Created  string
	Updated  string
}

type DeviceToSensor struct {
	Id  int
	DeviceID       int
	SensorID   int
	Created  string
	Updated  string
}

type DeviceToProtocol struct {
	Id  int
	DeviceID       int
	ProtocolID   int
	Created  string
	Updated  string
}

type DeviceToBuildingSystem struct {
	Id  int
	DeviceID int
	BuildingSystemID    int
	Created  string
	Updated  string
}