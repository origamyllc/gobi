package postgres

import (
	"testing"
)

func TestGetConnectionString(t *testing.T) {
	actual := GetConnectionString()
	expected := "user=postgres password=postgres dbname=light host=localhost port=5432 sslmode=disable";

	if actual != expected {
	t.Error("The connection string do not match")
	}
}

func TestInsert(t *testing.T) {
	db,err := Connect();
	if err != nil {
		t.Errorf("The Database %v failed to connect ",db)
	}

	query := `INSERT INTO users (age, email, first_name, last_name) VALUES (31, 'jon@cagghlhjjoun.io', 'Jonathan', 'Calhoun');`
	Insert(query,db)
}

func TestUpdate(t *testing.T) {
	db,err := Connect();
	if err != nil {
		t.Errorf("The Database %v failed to connect ",db)
	}

	query := `UPDATE users SET first_name = 'Johnny', last_name = 'Applesjeed' WHERE id = 8;`
	Update(query,db)
}

func TestDelete(t *testing.T) {
	db,err := Connect();
	if err != nil {
		t.Errorf("The Database %v failed to connect ",db)
	}

	query := `DELETE FROM users WHERE id = 1;`
	Delete(query,db)
}

func TestGet(t *testing.T) {
	db, err := Connect();
	if err != nil {
		t.Errorf("The Database %v failed to connect ",db)
	}

	type User struct {
		ID        int
		Age       int
		FirstName string
		LastName  string
		Email     string
	}

	sqlStatement := `SELECT * FROM users`
	var user User
     rows,err := Get(sqlStatement,db)
	for rows.Next() {
		error := rows.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email)
		if (error != nil) {
			t.Error("Can not get records ",db)
		}
	}
	defer rows.Close()
}

func TestConnect(t *testing.T) {
	db,err := Connect();
    if err != nil {
	  t.Errorf("The Database %v failed to connect",db)
    }
}


