package Observable

/*
	CREATE TABLE observable (
	observable_id                   integer PRIMARY KEY NOT NULL,
	observable_type                  varchar(50) NOT NULL,
	observable_name                  varchar(50) NOT NULL,
	observable_threshold_value                  varchar(50) NOT NULL,
	observable__unit                  varchar(50) NOT NULL
	)
*/

type Observable struct {
	ObservableID int
	ObservableName string
	ObservableType string
	ObservableThreholdValue string
	ObservableUnit string
}
