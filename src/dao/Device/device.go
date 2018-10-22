package Device

type Device struct {
	DeviceID   string
    DeviceName string
	DeviceMake string
	DeviceModel string
	DeviceType string
	DeviceIPAddress string
	DeviceMaccAddress string
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