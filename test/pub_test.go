package test

import (
	gorabbitmq "github.com/cr-mao/go-rabbitmq"
	"testing"
)

func TestProductDirect(t *testing.T) {

	var connParams = &gorabbitmq.ConnParams{
		User:     USER,
		Password: PASSWORD,
		Host:     HOST,
		Vhost:    VHOST,
		Port:     PORT,
	}

	mq, err := gorabbitmq.Conn(connParams)
	if err != nil {
		t.Fatalf("Failed to connect err:%v", err)
	}

	err = mq.ExchangeDeclare(EXCHANGE_NAME, "direct")
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
	var connParams = &gorabbitmq.ConnParams{
		User:     USER,
		Password: PASSWORD,
		Host:     HOST,
		Vhost:    VHOST,
		Port:     PORT,
	}
	mq, err := gorabbitmq.Conn(connParams)
	err = mq.ExchangeDeclare(EXCHANGE_NAME, "direct")
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
