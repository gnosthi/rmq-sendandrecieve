package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"github.com/gnosthi/rmq-sendandrecieve/rmq-snr-config"
	"log"
	"bytes"
	"time"
)

func failOnError(err error, msq string) {
	if err != nil {
		log.Fatal("%s: %s", msq, err)
		panic(fmt.Sprintf("%s : %s", msq, err))
	}
}

func main() {
	conn, err := amqp.Dial("amqp://"+rmq_snr_config.RmqUser+":"+rmq_snr_config.RmqPass+"@"+rmq_snr_config.RmqHost+":"+rmq_snr_config.RmqPort+"/")
	failOnError(err, "Failed to connect to RabbitMQ instance")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open channel.")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		rmq_snr_config.ChannelQueName, //name
		rmq_snr_config.RmqDurableMessage, //durable
		false, //autodelete
		false, //exclusive
		false, //noWait
		nil, //args
	)
	failOnError(err, "Failed to declare queue")

	err = ch.Qos(
		rmq_snr_config.RmqQOSCount, //count
		rmq_snr_config.RmqQOSSize,  //size
		rmq_snr_config.RmqQOSGlobal, //global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name,
		"",
		rmq_snr_config.RmqRecvAck, //acknowledge
		false, //exclusive
		false, //noLocal
		false, //noWait
		nil, //args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Recieved a message: %s", d.Body)
			if rmq_snr_config.RmqSendPersistMode == true {
				dot_count := bytes.Count(d.Body, []byte("."))
				t := time.Duration(dot_count)
				time.Sleep(t * time.Second)
				log.Printf("Done")
			} else {
				log.Printf("Done")
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
