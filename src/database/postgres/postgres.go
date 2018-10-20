package postgres
// ref --> https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.4.html/
import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"gobi/resouces"
	"log"
)

// Config holds the configuration used for instantiating a new Roach.
type Config struct {
	// Address that locates our postgres instance
	Host string
	// Port to connect to
	Port string
	// User that has access to the database
	User string
	// Password so that the user can login
	Password string
	// Database to connect to (must have been created priorly)
	Database string
}



func GetConnectionString() (string){
	// todo:= make env configurable and add thread pooling
	bytes := []byte(resouces.GetDatabaseConstants());
	var config  Config
	json.Unmarshal(bytes, &config)
	cfg := Config{Host:config.Host,Port:config.Port,User:config.User,Password:config.Password,Database:config.Database}
	conn_string := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.Database, cfg.Host, cfg.Port)
	    return conn_string
}

func Connect()( *sql.DB,error){
	psqlInfo := GetConnectionString();
	db, err :=  sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error Connecting to database"+psqlInfo)
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error Connecting to database"+psqlInfo)
		panic(err)
	}

	log.Println("Successfully connected to database!")
	log.Println(psqlInfo);
	return db,nil;
}

func Insert(query string,db *sql.DB)(error){
	log.Println("Executing Query := "+query)
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
		return err ;
	}
	return nil;
}

func Update(query string,db *sql.DB)(error){
	log.Println("Executing Query := "+query)
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
		return err ;
	}
	return nil ;
}

func Delete(query string,db *sql.DB)(error){
	log.Println("Executing Query := "+query)
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
		return err ;
	}
	return nil;
}

func Get(query string,db *sql.DB)(*sql.Rows,error){
	log.Println("Executing Query := "+query)
	rows,err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows,err
}

