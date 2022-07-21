package go_rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

/**
消费消息
queue 队列
Name 消费者名
callbak 回调函数
*/
func (this *MQ) Sub(queue string, Name string, callbak func(a <-chan amqp.Delivery, b string)) {
	msgs, err := this.Channel.Consume(queue, Name, false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	callbak(msgs, Name)
}

//消费者限流     消息没处理完 要阻塞 等待 前面的消息被ack 掉 才行   ，快速发送3条数据，发现只有2条消息内容，后面 那条要等 前面的消息被ack 后才会出来
//Qos(2, 0, false)
func (this *MQ) Qos(prefetchCount, prefetchSize int, global bool) error {
	return this.Channel.Qos(prefetchCount, prefetchSize, global)
}
