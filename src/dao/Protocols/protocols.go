package Protocols

/**
CREATE TABLE protocols (
   protocol_id                   integer PRIMARY KEY NOT NULL,
      protocol_name       varchar(256) NOT NULL,
      protocol_description       varchar(256) NOT NULL
)

 */

type Protocol struct {
	ProtocolID          int
	ProtocolName        string
	ProtocolDescription string
}
