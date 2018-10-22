package sql


/**
create table

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

CREATE TABLE tokens (
   user_id     integer PRIMARY KEY NOT NULL,
   token       varchar(256) NOT NULL
)

INSERT INTO users (username,  password, secret,first_name, last_name, email, organization_name )
VALUES ('username','password', 'secret' ,'first_name', 'last_name', 'email', 'organization_name' )
INSERT INTO tokens ( user_id,  token ) VALUES ('username','password')
 */

const GET_USERS_QUERY  = `SELECT * FROM users`;
const GET_USER_BY_ID_QUERY = "SELECT * FROM users WHERE user_id=";
const GET_USER_BY_USER_NAME_AND_PASSWORD_QUERY = "SELECT * FROM users WHERE username = '%s' AND password = '%s' ";
const ADD_USER_QUERY = "INSERT INTO users (username,  password, secret, first_name, last_name, email, organization_name ) VALUES ('%s','%s','%s','%s','%s','%s','%s');"
const DELETE_USER_BY_ID_QUERY = "DELETE FROM users WHERE user_id= "
const UPDATE_USER_NAME__QUERY =  `UPDATE users SET first_name = '%s', last_name = '%s' WHERE user_id= %d;`
const ADD_TOKEN_QUERY = "INSERT INTO tokens ( user_id,  token ) VALUES ('%d','%s')"
const GET_USER_TOKEN_QUERY = "SELECT token FROM tokens WHERE token='%s' "