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

create table gateway_to_protocol
    (
	  id                   integer PRIMARY KEY NOT NULL,
      gateway_id          integer REFERENCES gateway NOT NULL,
      protocol_id   integer  REFERENCES protocols NOT NULL
	)

create table gateway_to_device
    (
	  id                   integer PRIMARY KEY NOT NULL,
      gateway_id          integer REFERENCES gateway NOT NULL,
      device_id   integer  REFERENCES device NOT NULL
	)
alter table gateway  add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table gateway   add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table gateway_to_device  add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table gateway_to_device   add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table gateway_to_protocol  add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table gateway_to_protocol   add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();



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
	Created  string
	Updated  string
}

type GatewayToProtocol struct {
	Id int
	GatewayID int
	ProtocolID int
	Created  string
	Updated  string
}

type GatewayToDevice struct {
	Id int
	GatewayID int
	DeviceID int
	Created  string
	Updated  string
}

