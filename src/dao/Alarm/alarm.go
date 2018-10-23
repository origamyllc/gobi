package Alarm

/**
CREATE TABLE alarm (
	  id                   integer PRIMARY KEY NOT NULL,
      organization_id   integer  REFERENCES organization NOT NULL,
	  sub_organization_id   integer  REFERENCES sub_organization,
	  building_id   integer  REFERENCES building NOT NULL,
	  floor_id   integer  REFERENCES floor NOT NULL,
	  room_id   integer  REFERENCES room NOT NULL,
	  gateway_id   integer  REFERENCES gateway NOT NULL,
	  device_id   integer  REFERENCES device NOT NULL,
	  sensor_id   integer  REFERENCES sensor NOT NULL,
	  protocol_id   integer  REFERENCES protocols NOT NULL,
	  observable_id   integer  REFERENCES observable NOT NULL,
	  building_system_id   integer  REFERENCES building_system NOT NULL,
	  reading    varchar(50) NOT NULL,
	  reading_unit   varchar(50)  NOT NULL,
	  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)
*/

