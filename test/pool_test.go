package test

import (
	"fmt"
	gorabbitmq "github.com/cr-mao/go-rabbitmq"
	"log"
	"testing"
)

var globalPool *gorabbitmq.Pool

func init() {
	var connParams = &gorabbitmq.ConnParams{
		User:     USER,
		Password: PASSWORD,
		Host:     HOST,
		Vhost:    VHOST,
		Port:     PORT,
	}
	var err error
	globalPool, err = gorabbitmq.NewPool(10, connParams)
	if err != nil {
		log.Fatalf("Failed to create pool %v", err)
	}

	mq := globalPool.Get()
	err = mq.ExchangeDeclare(EXCHANGE_NAME, "direct")
	if err != nil {
		log.Fatalf("exchange声明失败 %+v", err)
	}
	//exchange绑定队列
	err = mq.DecQueueAndBind(QUEUE_NAME, BINGDING_KEY, EXCHANGE_NAME)
	if err != nil {
		log.Fatalf("exchange声明失败 %+v", err)
	}
}

func TestPool(t *testing.T) {

	mq1 := globalPool.Get()
	//fmt.Printf("%p\n", mq)
	mq2 := globalPool.Get()
	mq3 := globalPool.Get()
	mq4 := globalPool.Get()
	mq5 := globalPool.Get()
	mq6 := globalPool.Get()
	mq7 := globalPool.Get()
	mq8 := globalPool.Get()
	mq9 := globalPool.Get()
	mq10 := globalPool.Get()

	fmt.Printf("%p\n", globalPool.Get())
	fmt.Printf("%p\n", globalPool.Get())
	fmt.Printf("%p\n", globalPool.Get())
	fmt.Printf("%p\n", globalPool.Get())
	fmt.Printf("%p\n", globalPool.Get())
	fmt.Printf("%p\n", globalPool.Get())
	fmt.Printf("%p\n", globalPool.Get())
	fmt.Printf("%p\n", globalPool.Get())
	fmt.Printf("%p\n", globalPool.Get())
	fmt.Printf("%p\n", globalPool.Get())
	fmt.Printf("%p\n", globalPool.Get())
	fmt.Printf("%p\n", globalPool.Get())
	fmt.Printf("%p\n", globalPool.Get())
	fmt.Printf("%p\n", globalPool.Get())
	var err error

	err = mq1.SendMessage(ROUTING_KEY, EXCHANGE_NAME, "hello world 1")
	if err != nil {
		t.Log(err)
	}
	err = mq2.SendMessage(ROUTING_KEY, EXCHANGE_NAME, "hello world 2")
	if err != nil {
		t.Log(err)
	}

	err = mq3.SendMessage(ROUTING_KEY, EXCHANGE_NAME, "hello world 3")
	if err != nil {
		t.Log(err)
	}

	err = mq4.SendMessage(ROUTING_KEY, EXCHANGE_NAME, "hello world 4")
	if err != nil {
		t.Log(err)
	}
	err = mq5.SendMessage(ROUTING_KEY, EXCHANGE_NAME, "hello world 5")
	if err != nil {
		t.Log(err)
	}
	err = mq6.SendMessage(ROUTING_KEY, EXCHANGE_NAME, "hello world 6")
	if err != nil {
		t.Log(err)
	}
	err = mq7.SendMessage(ROUTING_KEY, EXCHANGE_NAME, "hello world 7")
	if err != nil {
		t.Log(err)
	}
	err = mq8.SendMessage(ROUTING_KEY, EXCHANGE_NAME, "hello world 8")
	if err != nil {
		t.Log(err)
	}
	err = mq9.SendMessage(ROUTING_KEY, EXCHANGE_NAME, "hello world 9")
	if err != nil {
		t.Log(err)
	}
	err = mq10.SendMessage(ROUTING_KEY, EXCHANGE_NAME, "hello world 10")
	if err != nil {
		t.Log(err)
	}
	err = mq1.SendMessage(ROUTING_KEY, EXCHANGE_NAME, "hello world 11")
	if err != nil {
		t.Log(err)
	}
}
