package memphis

import (
	"context"
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
	client, err := conn.CreateConsumer(
		stationName,
		consumerName,
		memphis.PullInterval(1*time.Second),
		memphis.ConsumerGroup(consumerGroup),
		memphis.BatchSize(30),
		memphis.MaxAckTime(15*time.Second),
		memphis.ConsumerGenUniqueSuffix(),
	)
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
			if e.Error() == "memphis: timeout" {
				return
			}
			err <- e
			return
		}
		for _, msg := range msgs {
			msgChan <- string(msg.Data())
			e = msg.Ack()
			if e != nil {
				err <- e
				return
			}
		}
	}
	e := c.client.Consume(handler)
	if e != nil {
		err <- e
	}
}
func (c *consumer) Close() error {
	// Not calling Destroy() because it will cause the consumer will receive all acknowledged messages again
	// https://discord.com/channels/963333392844328961/1074408130345177149

	//err := c.client.Destroy()
	//if err != nil {
	//	return err
	//}
	c.connection.Close()
	return nil
}
