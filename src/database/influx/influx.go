package main

import (
	"errors"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"time"
)


// queryDB convenience function to query the database
func queryDB(clnt client.Client, database_name string,cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database:  database_name,
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

func Connect()(client.Client,error) {
	// Create a new HTTPClient
	//todo:secure with username and password
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		log.Fatal(err)
	}
	return c,nil;
}

func CreateDatabase(c client.Client,database_name string)(string,error){
	created, err := queryDB(c, database_name,fmt.Sprintf("CREATE DATABASE %s", database_name))
	if err != nil {
		log.Fatal(err)
		return "",errors.New("Database Not Created")
	}
	log.Print(created)
	return database_name,nil;
}

func CreateNewBatchPoint(c client.Client,database_name string)(client.BatchPoints,error){
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database_name,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
		return nil,errors.New("Batchpoint Not Created")
	}
	return bp,nil
}

func AddPointToBatch(bp client.BatchPoints,tags map[string]string,fields map[string]interface{}, table_name string){
	pt, err := client.NewPoint(table_name, tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)
}

func WriteBatch(c client.Client,bp client.BatchPoints){
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
}

func Close(c client.Client){
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

func SELECT(c client.Client,query string,database_name string)([]interface{}){
	res, err := queryDB(c,database_name,query)
	 result := make([]interface{}, 1)
	//    s := make([]string, 3)
	if err != nil {
		log.Fatal(err)
	}

	for i, row := range res[0].Series[0].Values {
		result = append(result, row)
		log.Print(i, row)
	}

	return result;
}


func DELETE(c client.Client,query string,database_name string){
	res, err := queryDB(c,database_name,query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(res)
}


