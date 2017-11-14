package main

import (
	"github.com/streadway/amqp"
	"github.com/gnosthi/rmq-sendandrecieve/rmq-snr-config"
	"log"
	"fmt"
	"time"
	"io/ioutil"
	"strings"
	"math/rand"
)

func failOnError(err error, msq string) {
	if err != nil {
		log.Fatal("%s: %s", msq, err)
		panic(fmt.Sprintf("%s : %s", msq, err))
	}
}

func main() {
	conn, err := amqp.Dial("amqp://"+rmq_snr_config.RmqUser+":"+rmq_snr_config.RmqPass+"@"+rmq_snr_config.RmqHost+":"+rmq_snr_config.RmqPort+"/")
	failOnError(err, "Failed to Connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open channel")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		rmq_snr_config.ChannelQueName,
		false,
		false,
		false,
		false,
		nil,
	)

	body := ""
	if rmq_snr_config.MessageBody == "random" {
		rand.Seed(time.Now().Unix())
		content, err := ioutil.ReadFile("./quotes.txt")

		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}
		lines := strings.Split(string(content), "\n")
		quote := rand.Int() % len(lines)
		body = " " + lines[quote]
	} else {
		body = rmq_snr_config.MessageBody
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		})
		failOnError(err, "Failed to publish message")
}


