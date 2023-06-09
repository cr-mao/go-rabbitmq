package go_rabbitmq

import "sync/atomic"

type Pool struct {
	count uint64
	index uint64
	conns []*MQ
}

func NewPool(count int, connParams *ConnParams) (*Pool, error) {
	p := &Pool{count: uint64(count), conns: make([]*MQ, count)}
	for i := 0; i < count; i++ {
		conn, err := Conn(connParams)
		if err != nil {
			return nil, err
		}
		p.conns[i] = conn
	}
	return p, nil
}

func (p *Pool) Get() *MQ {
	return p.conns[int(atomic.AddUint64(&p.index, 1)%p.count)]
}
