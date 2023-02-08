package memphis

import (
	"github.com/memphisdev/memphis.go"
	"ww-api/pkg/queue"
)

type producer struct {
	client     *memphis.Producer
	connection *memphis.Conn
}

func NewProducer(user, token, memphisUrl, stationName, producerName string) (queue.Producer, error) {
	conn, err := memphis.Connect(memphisUrl, user, token)
	if err != nil {
		return nil, err
	}
	client, err := conn.CreateProducer(stationName, producerName)
	if err != nil {
		return nil, err
	}
	return &producer{
		connection: conn,
		client:     client,
	}, nil
}

func (p *producer) Publish(message []byte) error {
	return p.client.Produce(message)
}

func (p *producer) Close() error {
	p.connection.Close()
	return nil
}