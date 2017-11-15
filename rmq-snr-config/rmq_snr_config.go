package rmq_snr_config

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

var (

	RmqHost string
	RmqPort	string
	RmqUser	string
	RmqPass string

	ChannelQueName string
	MessageBody string

	config		*configStruct
)

var (
	RmqSendPersistMode bool
	RmqDurableMessage bool
)

var (
	RmqRecvAck bool
	RmqQOSCount int
	RmqQOSSize int
	RmqQOSGlobal bool
)

type configStruct struct {
	RmqHost		string  `json:"Host"`
	RmqPort		string	`json:"Port"`
	RmqUser 	string	`json:"User"`
	RmqPass		string	`json:"Pass"`
	ChannelQueName	string	`json:"QueueName"`
	MessageBody	string	`json:"MessageText"`
	RmqSendPersistMode bool `json:"SenderPersistMode"`
	RmqDurableMessage bool	`json:"DurableMessages"`
	RmqRecvAck	bool	`json:"RecieverAcknowledgeMessages"`
	RmqQOSCount int	`json:"QoSCount"`
	RmqQOSSize int	`json:"QoSSize"`
	RmqQOSGlobal bool `json:"QoSGlobal"`
}

func ReadConfig() error {
	fmt.Println("Reading from config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println("Error in reading config file ./config.json")
		panic(err)
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	RmqHost = config.RmqHost
	RmqPort	= config.RmqPort
	RmqUser = config.RmqUser
	RmqPass = config.RmqPass
	ChannelQueName = config.ChannelQueName
	MessageBody = config.MessageBody

	RmqSendPersistMode = config.RmqSendPersistMode
	RmqDurableMessage = config.RmqDurableMessage
	RmqRecvAck = config.RmqRecvAck
	RmqQOSCount = config.RmqQOSCount
	RmqQOSSize = config.RmqQOSSize
	RmqQOSGlobal = config.RmqQOSGlobal

	return nil
}