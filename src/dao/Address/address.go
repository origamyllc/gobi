package Address

/**
CREATE TABLE address (
   address_id                   integer PRIMARY KEY NOT NULL,
   address_line_1       varchar(256) NOT NULL,
   address_line_2       varchar(256) NOT NULL,
   city                 varchar(50) NOT NULL,
   state                varchar(50) NOT NULL,
   zip                  varchar(10) NOT NULL,
   country              varchar(10) NOT NULL,
   country_code         varchar(10) NOT NULL,
   lat                  varchar(10) NOT NULL,
   long                  varchar(10) NOT NULL
)

 */
type Address struct {
	ID           int
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	Zip          string
	Country      string
	CountryCode  string
	Lat          float64
	Long         float64
}