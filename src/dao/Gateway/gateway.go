package Gateway

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
}

type GatewayToProtocol struct {
	GatewayID int
	ProtocolID int
}

type GatewayToSupportedDeviceTypes struct {
	GatewayID int
	DeviceTypeID int
}

type GatewayToSupportedSensorTypes struct {
	GatewayID int
	SensorTypeID int
}

type GatewayToDevice struct {
	GatewayID int
	DeviceID int
}

type GatewayToSensor struct {
	GatewayID int
	SensorID int
}