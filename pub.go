package go_rabbitmq

import "github.com/streadway/amqp"

//生产消息
func (this *MQ) SendMessage(routing_key string, exchange string, message string) error {
	return this.Channel.Publish(exchange, routing_key, true, false,
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(message),
			DeliveryMode: 2,
		},
	)
}

//生成消息并开启 confirm 机制, confirmCallback 没收到处理函数
func (this *MQ) SendMessageWithConfirm(routing_key string, exchange string, message string, confirmCallback ...func()) error {
	this.SetConfirm() //开启confirm模式

	/**
	mandatory :
	 如果为true，在exchange正常且可到达的情况下。如果exchange+routeKey无法投递给queue，那么MQ会将消息返还给生产者;
	  如果为false时，则直接丢弃
	*/
	//this.NotifyReturn() //监听return
	err := this.Channel.Publish(exchange, routing_key, true, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	//判断是否收到
	if !this.ListenConfirm() {
		if len(confirmCallback) > 0 {
			confirmCallback[0]()
		}
	}
	return err
}

//发送延迟消息
func (this *MQ) SendDelayMessage(routing_key string, exchange string, message string, delay int) error {
	err := this.Channel.Publish(exchange, routing_key, true, false,
		amqp.Publishing{
			Headers:     map[string]interface{}{"x-delay": delay},
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	return err
}

// 发送延迟消息，并开启确认机制
func (this *MQ) SendDelayMessageWithConfirm(routing_key string, exchange string, message string, delay int) error {
	err := this.Channel.Publish(exchange, routing_key, true, false,
		amqp.Publishing{
			Headers:     map[string]interface{}{"x-delay": delay},
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	return err
}
