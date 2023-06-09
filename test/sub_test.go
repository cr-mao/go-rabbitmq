package test

import (
	"fmt"
	gorabbitmq "github.com/cr-mao/go-rabbitmq"
	"github.com/streadway/amqp"
	"testing"
)

func TestSub(t *testing.T) {
	var connParams = &gorabbitmq.ConnParams{
		User:     USER,
		Password: PASSWORD,
		Host:     HOST,
		Vhost:    VHOST,
		Port:     PORT,
	}

	mq, err := gorabbitmq.Conn(connParams)
	if err != nil {
		t.Fatalf("connection failed: %v", err)
	}

	//defer mq.CloseChannel()
	err = mq.ExchangeDeclare(EXCHANGE_NAME, "direct")
	if err != nil {
		t.Fatalf("exchange声明失败 %+v", err)
	}
	//exchange绑定队列
	err = mq.DecQueueAndBind(QUEUE_NAME, BINGDING_KEY, EXCHANGE_NAME)
	if err != nil {
		t.Fatalf("exchange声明失败 %+v", err)
	}
	mq.Sub(QUEUE_NAME, "consume1", sub)
	select {}
	defer mq.Close()
}

func sub(msgDelivery <-chan amqp.Delivery, b string) {
	for msg := range msgDelivery {
		fmt.Println(string(msg.Body))
		msg.Ack(false) //收到ask
	}
	//msg.Reject(false) //丢弃原消息
}
