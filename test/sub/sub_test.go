package sub

import (
	"fmt"
	gorabbitmq "github.com/cr-mao/go-rabbitmq"
	"github.com/streadway/amqp"
	"testing"
)

const (
	EXCHANGE_NAME = "mq_exchange_name_test"
	QUEUE_NAME    = "mq_queue_name_test"
	BINGDING_KEY  = "mq_bing_key_name_test"
)

func TestSub(t *testing.T) {
	mq := gorabbitmq.Conn("guest", "guest", "127.0.0.1", "/", 5672)
	//defer mq.CloseChannel()
	err := mq.ExchangeDeclare(EXCHANGE_NAME, "direct")
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
