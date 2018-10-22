package Floor

type Floor struct {
	FloorID                                 int
	FloorType                               int
	FloorName                               int
	FloorDescription                        string
	FloorTotlalNoOfPeople                   int
	FloorTotalArea                          int
	FloorAverageHourlyEnergyConsumption     float64
	FloorAverageWeeklyEnergyConsumption     float64
	FloorAverageMonthlyEnergyConsumption    float64
	FloorAverageYearlyEnergyConsumption     float64
}

type FloorToBuildingSystem struct {
	FloorID   int
	BuildingSystemID    int
}

type FloorToRoom struct {
	FloorID   int
	RoomID    int
}

type FloorToGateway struct {
	FloorID   int
	GatewayID    int
}