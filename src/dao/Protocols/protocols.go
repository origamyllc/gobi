package Protocols

/**
CREATE TABLE protocols (
   protocol_id                   integer PRIMARY KEY NOT NULL,
      protocol_name       varchar(256) NOT NULL,
      protocol_description       varchar(256) NOT NULL
)

alter table protocols   add column  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
alter table protocols    add column 	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();


 */

type Protocol struct {
	ProtocolID          int
	ProtocolName        string
	ProtocolDescription string
}
