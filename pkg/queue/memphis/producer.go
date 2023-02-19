package memphis

import (
	"github.com/memphisdev/memphis.go"
	"github.com/rs/zerolog/log"
	"ww-api/pkg/queue"
)

type producer struct {
	client     *memphis.Producer
	connection *memphis.Conn
}

func NewProducer(user, token, memphisUrl, stationName, producerName string) (queue.Producer, error) {
	conn, err := memphis.Connect(memphisUrl, user, token)
	if err != nil {
		log.Debug().Err(err).Msg("memphis producer connection error")
		return nil, err
	}
	client, err := conn.CreateProducer(
		stationName,
		producerName,
		memphis.ProducerGenUniqueSuffix(),
	)
	if err != nil {
		log.Debug().Err(err).Msg("memphis producer creation error")
		return nil, err
	}
	return &producer{
		connection: conn,
		client:     client,
	}, nil
}

func (p *producer) Publish(message []byte) error {
	return p.client.Produce(
		message,
		memphis.AckWaitSec(15),
		memphis.AsyncProduce(),
	)
}

func (p *producer) Close() error {
	err := p.client.Destroy()
	if err != nil {
		log.Debug().Err(err).Msg("memphis producer destroy error")
		return err
	}
	p.connection.Close()
	return nil
}
