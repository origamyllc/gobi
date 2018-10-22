package Sensor

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