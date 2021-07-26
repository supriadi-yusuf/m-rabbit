package mrabbit

import (
	"time"

	"github.com/streadway/amqp"
)

const (
	MaxTrial1 = 3
	MaxTrial2 = 7
)

type RealRabbit struct {
	NotifyListener
	conn *amqp.Connection
	ch   *amqp.Channel
}

func CreateRealRabbitInstance() IRabbit {
	realRabbit := new(RealRabbit)
	//realRabbit.GetNotifyCloseChannel()
	return realRabbit
}

func (q *RealRabbit) DialRabbit(uri string, commandNo int) error {
	var err error
	q.conn, err = amqp.Dial(uri)
	return err
}

func (q *RealRabbit) CreateChannel() error {
	var err error
	q.ch, err = q.conn.Channel()
	return err
}

func (q *RealRabbit) CloseChannel() error {
	return q.ch.Close()
}

func (q *RealRabbit) CloseConnection() error {
	return q.conn.Close()
}

func (q *RealRabbit) SetHeartBeat(duration time.Duration) {
	q.conn.Config.Heartbeat = duration
}

func (q *RealRabbit) DeclareQueue(queue string, durable, autodelete, exclusive, nowait bool, args amqp.Table) (amqp.Queue, error) {
	return q.ch.QueueDeclare(queue, durable, autodelete, exclusive, nowait, args)
}

func (q *RealRabbit) NotifyClose(ch chan *amqp.Error) chan *amqp.Error {
	q.NotifyCloseCh = ch
	return q.conn.NotifyClose(ch)
}

func (q *RealRabbit) ChannelQos(prefetchCount, prefetchSize int, global bool) error {
	return q.ch.Qos(prefetchCount, prefetchSize, global)
}

func (q *RealRabbit) ConsumeQueueMsq(queue, consumer string, autoAck, exclusive,
	noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	return q.ch.Consume(queue, consumer, autoAck, exclusive,
		noLocal, noWait, args)
}

func (q *RealRabbit) GetWaitingTimeToConnect(trial int) time.Duration {
	switch {
	case trial <= MaxTrial1:
		return time.Duration(30) * time.Second

	case trial <= MaxTrial2:
		return time.Duration(10) * time.Minute

	default:
		return time.Duration(1) * time.Hour
	}
}
