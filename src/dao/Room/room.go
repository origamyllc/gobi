package Room

type Room struct {
	RoomID                                 int
	RoomType                               int
	RoomName                               int
	RoomDescription                        string
	RoomTotalArea                          int
	RoomAverageHourlyEnergyConsumption     float64
	RoomAverageWeeklyEnergyConsumption     float64
	RoomAverageMonthlyEnergyConsumption    float64
	RoomAverageYearlyEnergyConsumption     float64
}

type RoomToGateway struct {
	RoomID   int
	GatewayID    int
}