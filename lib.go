package go_rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strings"
)

//mq 对象
var mq *MQ

type MQ struct {
	Conn          *amqp.Connection
	Channel       *amqp.Channel
	notifyConfirm chan amqp.Confirmation
	notifyReturn  chan amqp.Return
}

//连接返回 MQ对象，已经初始化连接，和 amqp.Channel
func Conn(user, password, host, vhost string, port int) *MQ {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d%s", user, password, host, port, vhost)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Fatal(err)
	}
	mq = &MQ{
		Conn: conn,
	}
	channel, err := mq.Conn.Channel()
	if err != nil {
		panic(err)
	}
	mq.Channel = channel
	return mq
}

//关闭通道
func (this *MQ) CloseChannel() error {
	return this.Channel.Close()
}

func (this *MQ) Close() error {
	err := this.Channel.Close()
	if err != nil {
		return err
	}
	return this.Conn.Close()
}

//开启消息确认
func (this *MQ) SetConfirm() {
	err := this.Channel.Confirm(false)
	if err != nil {
		log.Println(err)
	}
	this.notifyConfirm = this.Channel.NotifyPublish(make(chan amqp.Confirmation))
}

//消息确认
func (this *MQ) ListenConfirm() bool {
	defer this.Channel.Close()
	ret := <-this.notifyConfirm
	return ret.Ack
}

/**
		mandatory :
		 如果为true，在exchange正常且可到达的情况下。如果exchange+routeKey无法投递给queue，那么MQ会将消息返还给生产者;
		如果为false时，则直接丢弃
exchange 和 routeKey 如果 发送 没有找到队列的情况 会往 chan amqp.Return 投递消息 ，我们就能捕获到了。
*/
func (this *MQ) NotifyReturn() {
	this.notifyReturn = this.Channel.NotifyReturn(make(chan amqp.Return))
	go this.listenReturn() //使用协程执行
}

func (this *MQ) listenReturn() {
	<-this.notifyReturn
	//ret := <-this.notifyReturn
	//if string(ret.Body) != "" {
	//	log.Println("消息没有正确入列:", string(ret.Body))
	//}
}

//申明队列以及绑定路由key,多个队列 可以用逗号分隔
func (this *MQ) DecQueueAndBind(queues string, bingding_key string, exchange string) error {
	qList := strings.Split(queues, ",")
	for _, queue := range qList {
		q, err := this.Channel.QueueDeclare(queue, false, false, false, false, nil)
		if err != nil {
			return err
		}
		err = this.Channel.QueueBind(q.Name, bingding_key, exchange, false, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

//声明exhchange
func (this *MQ) ExchangeDeclare(exchange_name, kind string) error {
	err := this.Channel.ExchangeDeclare(exchange_name, kind, false, false, false, false, nil)
	return err
}
