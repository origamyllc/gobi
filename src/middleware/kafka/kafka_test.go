package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"testing"
)

 func TestConnectToProducer(t *testing.T) {
	 actual,err := ConnectToProducer();
	 if err != nil {
		 panic(err)
	 }
	 cfg := &kafka.ConfigMap{"bootstrap.servers": "localhost:9092" }
	 expected, err := kafka.NewProducer(cfg)
	 if err != nil {
		 panic(err)
	 }

	 if actual == expected {
	 	fmt.Print(actual, expected)
		 t.Error("The connection string do not match")
	 }
 }

func TestPublish(t *testing.T) {
	cfg := &kafka.ConfigMap{"bootstrap.servers": "localhost:9092" }
	expected, err := kafka.NewProducer(cfg)
	if err != nil {
		panic(err)
	}
	Publish(expected, "find-device","Shize!!" )
	expected.Flush(20 * 1000)
	defer expected.Close()

}

func xTestConnectToConsumer(t *testing.T) {
	cfg := &kafka.ConfigMap{"bootstrap.servers": "localhost:9092" }
	p, err := kafka.NewProducer(cfg)
	if err != nil {
		panic(err)
	}
	Publish(p, "find-device","Shize!!" )
    x := Next("find-device")
    print(x)
	p.Flush(20 * 1000)
	defer p.Close()

}