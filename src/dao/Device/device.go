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

alter table device
      add constraint device_to_sensor
      foreign key (device_id)
      references sensor (sensor_id);

alter table device
      add constraint device_to_protocol
      foreign key (device_id)
      references protocols (protocol_id);

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
}

type DeviceToSensor struct {
	DeviceID       int
	SensorID   int
}

type DeviceToProtocol struct {
	DeviceID       int
	ProtocolID   int
}