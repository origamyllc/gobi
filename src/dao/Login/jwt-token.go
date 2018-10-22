package Login


/**

CREATE TABLE tokens (
   user_id     integer PRIMARY KEY NOT NULL,
   token       varchar(256) NOT NULL
)

INSERT INTO tokens ( user_id,  token ) VALUES ('username','password')
 */

type JwtToken struct {
	Token string `json:"token"`
}