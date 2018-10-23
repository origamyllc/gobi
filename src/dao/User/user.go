package User


/*
    CREATE TABLE users (
		user_id SERIAL NOT NULL PRIMARY KEY,
		username       varchar(40) NOT NULL,
		password       varchar(40) NOT NULL,
		secret         varchar(40) NOT NULL,
		first_name     varchar(40) NOT NULL,
		last_name      varchar(40) NOT NULL,
		email          varchar(100) NOT NULL,
		organization_name varchar (100) NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		UNIQUE(username),
		UNIQUE(password)
     );

	INSERT INTO users (username,  password, secret,first_name, last_name, email, organization_name )
	VALUES ('username','password', 'secret' ,'first_name', 'last_name', 'email', 'organization_name' )

*/

type User struct {
	ID        int               `json:"user_id"`
	Username     string         `json:"Username"`
	Password     string         `json:"password"`
	Secret       string         `json:"secret"`
	FirstName string            `json:"FirstName"`
	LastName  string            `json:"LastName"`
	Email     string            `json:"Email"`
	OrganizationName  string    `json:"OrganizationName"`
	Created  string            `json:"Created"`
	Updated  string            `json:"Updated"`
}

type UserResponse struct {
	ID        int               `json:"user_id"`
	Username     string         `json:"-"`
	Password     string         `json:"-"`
	Secret       string         `json:"secret"`
	FirstName string            `json:"FirstName"`
	LastName  string            `json:"LastName"`
	Email     string            `json:"Email"`
	OrganizationName  string    `json:"OrganizationName"`
	Created  string            `json:"Created"`
	Updated  string            `json:"Updated"`
}

