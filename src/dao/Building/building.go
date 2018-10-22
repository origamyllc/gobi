package Building

type Building struct {
	BuildingID                    int
	BuildingType                  string
	BuildingName                  string
	BuildingAddressID             int
	BuildingYearBuilt             string
	BuildingDescription           string
	BuildingTotlalNoOfPeople      int
	BuildingTotalArea             int
	BuildingTotalNoOfFloors       int
	BuildingHoursOccupied         string
	BuildingAverageHourlyEnergyConsumption     float64
	BuildingAverageWeeklyEnergyConsumption     float64
	BuildingAverageMonthlyEnergyConsumption    float64
	BuildingAverageYearlyEnergyConsumption     float64
}

type BuildingToFloor struct{
	BuildingID int
	FloorID    int
}

type BuildingToBuildingSystem struct {
	BuildingID int
	BuildingSystemID    int
}