package Room

/*
CREATE TABLE room (
    room_id                   integer PRIMARY KEY NOT NULL,
	room_type                  varchar(50) NOT NULL,
	room_name                  varchar(50) NOT NULL,
	room_description           varchar(256) NOT NULL,
	room_hours_occupied         varchar(50),
	room_average_hourly_energy_consumption     FLOAT(10),
	room_average_weekly_energy_consumption     FLOAT(10),
	room_average_monthly_energy_consumption    FLOAT(10),
	room_average_yearly_energy_consumption     FLOAT(10)
)

create table room_to_gateway
    (
	  id                   integer PRIMARY KEY NOT NULL,
      room_id          integer REFERENCES room NOT NULL,
      gateway_id   i   integer  REFERENCES gateway NOT NULL
	)

 */
type Room struct {
	RoomID                                 int
	RoomType                               int
	RoomName                               int
	RoomDescription                        string
	RoomTotalArea                          int
	RoomHoursOccupied                      string
	RoomAverageHourlyEnergyConsumption     float64
	RoomAverageWeeklyEnergyConsumption     float64
	RoomAverageMonthlyEnergyConsumption    float64
	RoomAverageYearlyEnergyConsumption     float64
}

type RoomToGateway struct {
	RoomID   int
	GatewayID    int
}