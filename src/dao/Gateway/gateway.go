package Gateway
/*
CREATE TABLE gateway (
	gateway_id                   integer PRIMARY KEY NOT NULL,
	gateway_type                  varchar(50) NOT NULL,
	gateway_name                  varchar(50) NOT NULL,
	gateway_make                  varchar(50) NOT NULL,
	gateway_model                  varchar(50) NOT NULL,
	gateway_ip_address             varchar(50) NOT NULL,
	gateway_mac_addess              varchar(50) NOT NULL,
	gateway_location              varchar(50) NOT NULL,
	gateway_description              varchar(256) NOT NULL,
	gateway_data_sheet_location       varchar(512) NOT NULL,
	gateway_average_hourly_energy_consumption     FLOAT(10),
	gateway_average_weekly_energy_consumption     FLOAT(10),
	gateway_average_monthly_energy_consumption    FLOAT(10),
	gateway_average_yearly_energy_consumption     FLOAT(10)
)

alter table gateway
      add constraint gateway_to_protocol
      foreign key (gateway_id)
      references protocols (protocol_id);

alter table gateway
      add constraint gateway_to_device
      foreign key (gateway_id)
      references device (device_id);


*/
type Gateway struct {
	GatewayID int
	GatewayName string
	GatewayMake string
	GatewayModel string
	GatewayType  string
	GatewayIP    string
	GatewayMACAddress string
	GatewayLocation string
	GatewayDatasheetLocation string
	GatewayAverageHourlyEnergyConsumption     float64
	GatewayAverageWeeklyEnergyConsumption     float64
	GatewayAverageMonthlyEnergyConsumption     float64
	GatewayAverageYearlyEnergyConsumption     float64
}

type GatewayToProtocol struct {
	GatewayID int
	ProtocolID int
}

type GatewayToDevice struct {
	GatewayID int
	DeviceID int
}

