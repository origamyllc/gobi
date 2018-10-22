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


alter table sensor
      add constraint sensor_to_protocol
      foreign key (sensor_id)
      references protocols (protocol_id);

alter table sensor
      add constraint sensor_to_observable
      foreign key (sensor_id)
      references observable (observable_id);
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