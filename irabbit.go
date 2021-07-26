package mrabbit

import (
	"time"

	"github.com/streadway/amqp"
)

type IRabbit interface {
	DialRabbit(string, int) error
	CreateChannel() error
	CloseChannel() error
	CloseConnection() error
	SetHeartBeat(time.Duration)
	DeclareQueue(string, bool, bool, bool, bool, amqp.Table) (amqp.Queue, error)
	NotifyClose(chan *amqp.Error) chan *amqp.Error
	ChannelQos(prefetchCount, prefetchSize int, global bool) error
	ConsumeQueueMsq(queue, consumer string, autoAck, exclusive, noLocal,
		noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
	GetNotifyCloseChannel() chan *amqp.Error
	GetWaitingTimeToConnect(int) time.Duration
}
