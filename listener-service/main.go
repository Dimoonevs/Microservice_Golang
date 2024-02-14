package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/imoonevs/icroservice_olang/listener-service/event"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// connect to rabbit mq
	rabbitConn, err := connction()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	// start to listen for messages
	log.Println("Listening for and consuming RabbitMQ messages...")
	// create consumer
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	// watch queue and comsum event
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connction() (*amqp.Connection, error) {
	var count int64
	var backOff = 1 * time.Second
	var conn *amqp.Connection

	// don't continue until rabbit is ready

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			count++
		} else {
			log.Println("Connceted to rabbit mq")
			conn = c
			break
		}
		if count > 5 {
			fmt.Println(err)
			return nil, err
		}
		backOff = time.Duration(math.Pow(float64(count), 2)) * time.Second
		fmt.Println("backing off...")
		time.Sleep(backOff)
		continue
	}

	return conn, nil
}
