package main

import (
	"log"
	"testing"
)

func TestGetConnectionString(t *testing.T) {
	Connect()
}

//
func TestCreateDatabase(t *testing.T) {
	c,_ := 	Connect()
	CreateDatabase(c,"square_holes")
	Close(c)
}

func TestCreateNewBatchPointatch(t *testing.T) {
	c,_ := 	Connect()
	database,err := CreateDatabase(c,"square_holes")
	if err != nil {
		panic(err)
	}
	CreateNewBatchPoint(c,database)
	Close(c)
}


func TestAddPointToBatch(t *testing.T) {
	c,_ := 	Connect()
	database,err := CreateDatabase(c,"square_holes")
	if err != nil {
		panic(err)
	}
	bp,err := CreateNewBatchPoint(c,database)
	if err != nil {
		panic(err)
	}
	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}

	AddPointToBatch(bp,tags,fields,"cpu_usage")
	Close(c)
}

func TestWriteBatch(t *testing.T) {
	c,_ := 	Connect()
	database,err := CreateDatabase(c,"square_holes")
	if err != nil {
		panic(err)
	}
	bp,err := CreateNewBatchPoint(c,database)
	if err != nil {
		panic(err)
	}
	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}

	AddPointToBatch(bp,tags,fields,"cpu_usage")
	WriteBatch(c ,bp);
	Close(c)
}

func TestSelect(t *testing.T) {
	c,_ := 	Connect()
	database,err := CreateDatabase(c,"square_holes")
	if err != nil {
		panic(err)
	}
	result := SELECT(c ,"SELECT * FROM cpu_usage",database)
	log.Println(result[1])

	//close(result)
	Close(c)
}

func TestDelete(t *testing.T) {
	c,_ := 	Connect()
	database,err := CreateDatabase(c,"square_holes")
	if err != nil {
		panic(err)
	}
	 DELETE(c ,`DELETE  FROM "cpu_usage" `,database)

	//close(result)
	Close(c)
}

