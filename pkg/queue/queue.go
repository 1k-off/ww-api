package queue

type Producer interface {
	Publish(message []byte) error
	Close() error
}

type Consumer interface {
	Consume(msgChan chan string, err chan error)
	Close() error
}
