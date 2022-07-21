package go_rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type MQconn struct {
	conn *amqp.Connection
}

var mqconn *MQconn

func Conn(user, password, host, vhost string, port int) *MQconn {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d%s", user, password, host, port, vhost)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Fatal(err)
	}
	mqconn = &MQconn{}
	mqconn.conn = conn
	return mqconn
}
