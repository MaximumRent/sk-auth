package worker

import (
	"github.com/streadway/amqp"
	"log"
	"sk-auth/errors"
)

type RabbitMqBroker struct {
	Broker
	ip         string
	port       string
	poolSize   int
	stop       chan bool
	connection *amqp.Connection
}

const (
	_AUTH_QUEUE_NAME = "auth"
)

func (broker *RabbitMqBroker) Start() {
	broker.stop = make(chan bool)
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		errors.FailOnError(err, "Error on RabbitMQ connecting.")
	}
	defer connection.Close()
	broker.connection = connection
	initQueue(connection)
	broker.receiveMessages(connection)
	<-broker.stop
}

func (broker *RabbitMqBroker) receiveMessages(connection *amqp.Connection) {
	for i := 0; i < broker.poolSize; i++ {
		channel, err := connection.Channel()
		errors.FailOnError(err, "Error on creating RabbitMQ channel.")
		go subscribeToQueue(channel)
	}
}

func subscribeToQueue(channel *amqp.Channel) {
	messages, err := channel.Consume(
		_AUTH_QUEUE_NAME, // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	errors.FailOnError(err, "Failed to register a consumer")
	for deliver := range messages {
		body := string(deliver.Body)

	}
}

func initQueue(connection *amqp.Connection) {
	channel, err := connection.Channel()
	errors.FailOnError(err, "Failed to open a channel")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		_AUTH_QUEUE_NAME, // name
		false,            // durable
		false,            // delete when usused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	errors.FailOnError(err, "Failed to declare a queue")
	purged, err := channel.QueuePurge(_AUTH_QUEUE_NAME, false)
	errors.FailOnError(err, "Failed on queue purging!")
	log.Printf("Purged %d messages from queue '%s'", purged, queue.Name)
}

func (broker *RabbitMqBroker) Stop() {
	broker.stop <- true
}

func getRabbitMqBroker(brokerDef map[interface{}]interface{}) *RabbitMqBroker {
	broker := new(RabbitMqBroker)
	broker.ip = brokerDef["ip"].(string)
	broker.port = brokerDef["port"].(string)
	broker.poolSize = brokerDef["poolSize"].(int)
	return broker
}