package worker

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"sk-auth/errors"
	"sk-auth/mongo"
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
	_AUTH_QUEUE_NAME = "sk-auth"
	_AUTH_EXCHANGE_NAME = "sk-auth-exchange"
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
	forever := make(chan bool)
	go func() {
		for deliver := range messages {
			body := string(deliver.Body)
			requestMessage := new(AuthRequestMessage)
			err := json.Unmarshal([]byte(body), requestMessage)
			if err != nil {
				log.Println("Can't unmarshal message from RabbitMQ. Cause: ", err)
			}
			processMessage(channel, deliver, requestMessage)
		}
	}()
	<-forever
}

func processMessage(channel *amqp.Channel, deliver amqp.Delivery,requestMessage *AuthRequestMessage) {

	err := mongo.ValidateAuthToken(requestMessage.Email, requestMessage.Nickname, requestMessage.Token)

	response := new(AuthResponseMessage)
	response.ReturnedMessage = requestMessage
	if err != nil {
		response.HasAccess = false
	} else {
		if err != nil {
			response.HasAccess = false
		} else {
			response.HasAccess = true
		}
	}
	jsonBytes, err := json.Marshal(response)

	err = channel.Publish(
		_AUTH_EXCHANGE_NAME,        // exchange
		deliver.ReplyTo, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: deliver.CorrelationId,
			Body:          jsonBytes,
		})

	deliver.Ack(false)
}

func initQueue(connection *amqp.Connection) {

	channel, err := connection.Channel()
	errors.FailOnError(err, "Failed to open a channel")
	defer channel.Close()

	err = channel.ExchangeDeclare(
		_AUTH_EXCHANGE_NAME,   // name
		"topic", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	channel.QueueBind(_AUTH_QUEUE_NAME, "", _AUTH_EXCHANGE_NAME, false, nil)
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
	broker.ip = brokerDef["address"].(string)
	broker.port = string(brokerDef["port"].(int))
	broker.poolSize = brokerDef["poolSize"].(int)
	return broker
}
