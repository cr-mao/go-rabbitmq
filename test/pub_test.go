package test

import (
	gorabbitmq "github.com/cr-mao/go-rabbitmq"
	"testing"
)

const (
	EXCHANGE_NAME = "mq_exchange_name_test"
	QUEUE_NAME    = "mq_queue_name_test"
	BINGDING_KEY  = "mq_bing_key_name_test"
	ROUTING_KEY   = "mq_bing_key_name_test"
)

func TestProductDirect(t *testing.T) {
	mq := gorabbitmq.Conn("guest", "guest", "127.0.0.1", "/", 5672)
	err := mq.ExchangeDeclare(EXCHANGE_NAME, "direct")
	if err != nil {
		t.Fatalf("exchange声明失败 %+v", err)
	}
	//exchange绑定队列
	err = mq.DecQueueAndBind(QUEUE_NAME, BINGDING_KEY, EXCHANGE_NAME)
	if err != nil {
		t.Fatalf("exchange声明失败 %+v", err)
	}

	err = mq.SendMessage(ROUTING_KEY, EXCHANGE_NAME, "hello world")
	if err != nil {
		t.Log(err)
	}

	mq.Close()
}

func TestProductDirectWithConfirm(t *testing.T) {
	mq := gorabbitmq.Conn("guest", "guest", "127.0.0.1", "/", 5672)
	err := mq.ExchangeDeclare(EXCHANGE_NAME, "direct")
	if err != nil {
		t.Fatalf("exchange声明失败 %+v", err)
	}
	//exchange绑定队列
	err = mq.DecQueueAndBind(QUEUE_NAME, BINGDING_KEY, EXCHANGE_NAME)
	if err != nil {
		t.Fatalf("exchange声明失败 %+v", err)
	}
	err = mq.SendMessageWithConfirm(ROUTING_KEY, EXCHANGE_NAME, "hello world")
	if err != nil {
		t.Log(err)
	}
	mq.Close()
}
