package mrabbit

import "github.com/streadway/amqp"

type NotifyListener struct {
	NotifyCloseCh chan *amqp.Error
}

func (listener *NotifyListener) GetNotifyCloseChannel() chan *amqp.Error {
	return listener.NotifyCloseCh
}
