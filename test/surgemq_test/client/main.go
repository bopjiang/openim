package main

import (
	"github.com/influxdata/surgemq/service"
	"github.com/surgemq/message"
	"log"
	"time"
)

func onPublish(msg *message.PublishMessage) error {
	log.Printf("push messsage to topic[%s], already subscribed, %s", string(msg.Topic()), msg)
	return nil
}

func main() {
	var err error

	log.Print("starting...\n")
	// Instantiates a new Client
	c := &service.Client{}

	// Creates a new MQTT CONNECT message and sets the proper parameters
	msg := message.NewConnectMessage()
	msg.SetWillQos(1)
	msg.SetVersion(4)
	msg.SetCleanSession(true)
	msg.SetClientId([]byte("surgemq"))
	msg.SetKeepAlive(10)
	msg.SetWillTopic([]byte("will"))
	msg.SetWillMessage([]byte("send me home"))
	msg.SetUsername([]byte("surgemq"))
	msg.SetPassword([]byte("verysecret"))

	host := "tcp://127.0.0.1:1883"
	// Connects to the remote server at 127.0.0.1 port 1883
	err = c.Connect(host, msg)
	if err != nil {
		log.Printf("failed to Connect, %s", err)
		return
	}

	// Creates a new SUBSCRIBE message to subscribe to topic "abc"
	submsg := message.NewSubscribeMessage()
	submsg.AddTopic([]byte("abc"), 0)

	// Subscribes to the topic by sending the message. The first nil in the function
	// call is a OnCompleteFunc that should handle the SUBACK message from the server.
	// Nil means we are ignoring the SUBACK messages. The second nil should be a
	// OnPublishFunc that handles any messages send to the client because of this
	// subscription. Nil means we are ignoring any PUBLISH messages for this topic.
	err = c.Subscribe(submsg, nil, onPublish)
	if err != nil {
		log.Printf("failed to Subscribe, %s", err)
		return
	}

	pktid := uint16(123)
	qos := byte(0) //TODO: can not receive PUBLISH when set to 1

	// Creates a new PUBLISH message with the appropriate contents for publishing
	pubmsg := message.NewPublishMessage()
	pubmsg.SetPacketId(pktid)
	pubmsg.SetTopic([]byte("abc"))
	pubmsg.SetPayload(make([]byte, 1024))
	pubmsg.SetQoS(qos)

	// Publishes to the server by sending the message
	err = c.Publish(pubmsg, nil)
	if err != nil {
		log.Printf("failed to Publish, %s, %s", msg, err)
		return
	}

	time.Sleep(time.Second)
	// Disconnects from the server
	c.Disconnect()

	log.Print("exit...\n")
}
