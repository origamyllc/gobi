package main

/**
Needs  kafka and rdkafka installation begore go get github.com/confluentinc/confluent-kafka-go/kafka
 */

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gobi/resouces"
)

type ProducerConfig struct {
	// Address that locates our kafka instance
	HOST string
	// port that locates our kafka instance
	PORT string
}

type ConsumerConfig struct {
	// Address that locates our kafka instance
	SERVERS string
	// port that locates our kafka instance
	GROUP_ID string

	//place to which offset gets reset
	AUTO_OFFSET_RESET string
}



func ConnectToProducer()(*kafka.Producer,error){
	// load the config from external lib
	// todo:= make env configurable and add thread pooling
	bytes := []byte(resouces.GetKafkaProducerConstants());
	var config  ProducerConfig
	json.Unmarshal(bytes, &config)
	cfg := &kafka.ConfigMap{"bootstrap.servers": config.HOST + ":"+ config.PORT }
	p, err := kafka.NewProducer(cfg)
	if err != nil {
		panic(err)
		return nil,errors.New("Kafka Connection to the producer failed")
	}
	fmt.Print("connection successful")
	return p,nil;
}

func Publish(p *kafka.Producer, topic_name string ,message string ){
	// Produce messages to topic (asynchronously)

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	topic := topic_name;
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(message),
		}, nil)
	// Wait for message deliveries before shutting down
}


func ConnectToConsumer()(*kafka.Consumer,error){
	bytes := []byte(resouces.GetKafkaConsumerConstants());
	var config  ConsumerConfig
	json.Unmarshal(bytes, &config)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.SERVERS,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
		return nil,errors.New("Kafka Connection to the consumer failed")
	}
	fmt.Println("connection successful")
	return c,nil;
}

func getSubscriber(c *kafka.Consumer,topic_name string)(chan string){
	ch := make(chan string)
	c.SubscribeTopics([]string{topic_name, "^aRegex.*[Tt]opic"}, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
            str := string(msg.Value)
			ch <- str
			fmt.Println(str)
			return ch;
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}


}

func Next(topic_name string)(string){
	c,err:= ConnectToConsumer()
	if err != nil {
		panic(err)
	}
	ch := getSubscriber(c,topic_name)
	x := <- ch;
	defer close(ch)
	defer c.Close()
    return x;
}




