package memphis

import (
	"context"
	"fmt"
	"github.com/memphisdev/memphis.go"
	"time"
	"ww-api/pkg/queue"
)

type consumer struct {
	client     *memphis.Consumer
	connection *memphis.Conn
}

func NewConsumer(user, token, memphisUrl, stationName, consumerName, consumerGroup string, ctx context.Context) (queue.Consumer, error) {
	conn, err := memphis.Connect(memphisUrl, user, token, memphis.Timeout(120*time.Second))
	if err != nil {
		return nil, err
	}
	client, err := conn.CreateConsumer(stationName, consumerName, memphis.PullInterval(1*time.Second), memphis.ConsumerGroup(consumerGroup))
	if err != nil {
		return nil, err
	}
	client.SetContext(ctx)
	return &consumer{
		connection: conn,
		client:     client,
	}, nil
}

func (c *consumer) Consume(msgChan chan string, err chan error) {
	handler := func(msgs []*memphis.Msg, e error, ctx context.Context) {
		if e != nil {
			fmt.Printf("Fetch failed: %v", err)
			return
		}
		for _, msg := range msgs {
			msgChan <- string(msg.Data())
			e = msg.Ack()
			if e != nil {
				err <- e
			}
		}
	}
	e := c.client.Consume(handler)
	if e != nil {
		err <- e
	}
}
func (c *consumer) Close() {
	c.connection.Close()
}
