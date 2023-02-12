package app

import (
	"encoding/json"
	"log"
	"time"
	"ww-api/pkg/entities"
	"ww-api/pkg/queue"
	"ww-api/pkg/queue/memphis"
)

type Manager struct {
	svc                             *Service
	sslDataProducer                 queue.Producer
	uptimeDataProducer              queue.Producer
	domainExpirationDataProducer    queue.Producer
	sslMetricsConsumer              queue.Consumer
	uptimeMetricsConsumer           queue.Consumer
	domainExpirationMetricsConsumer queue.Consumer
}

const (
	baseName = "api"
)

const (
	timePeriodSslTargets              = time.Minute * 5
	timePeriodUptimeTargets           = time.Minute * 1
	timePeriodDomainExpirationTargets = time.Minute * 1
)

func (s *Service) NewManager(mUser, mToken, mUrl, sslTSN, uptimeTSN, domainExpirationTSN, sslMSN, uptimeMSN, domainExpirationMSN string) (*Manager, error) {
	producerName := baseName
	consumerName := baseName
	consumerGroup := baseName
	sslDataProducer, err := memphis.NewProducer(mUser, mToken, mUrl, sslTSN, producerName)
	if err != nil {
		return nil, err
	}
	uptimeDataProducer, err := memphis.NewProducer(mUser, mToken, mUrl, uptimeTSN, producerName)
	if err != nil {
		return nil, err
	}
	domainExpirationDataProducer, err := memphis.NewProducer(mUser, mToken, mUrl, domainExpirationTSN, producerName)
	if err != nil {
		return nil, err
	}
	sslMetricsConsumer, err := memphis.NewConsumer(mUser, mToken, mUrl, sslMSN, consumerName, consumerGroup, s.ctx)
	if err != nil {
		return nil, err
	}
	uptimeMetricsConsumer, err := memphis.NewConsumer(mUser, mToken, mUrl, uptimeMSN, consumerName, consumerGroup, s.ctx)
	if err != nil {
		return nil, err
	}
	domainExpirationMetricsConsumer, err := memphis.NewConsumer(mUser, mToken, mUrl, domainExpirationMSN, consumerName, consumerGroup, s.ctx)
	if err != nil {
		return nil, err
	}
	return &Manager{
			svc:                             s,
			sslDataProducer:                 sslDataProducer,
			uptimeDataProducer:              uptimeDataProducer,
			domainExpirationDataProducer:    domainExpirationDataProducer,
			sslMetricsConsumer:              sslMetricsConsumer,
			uptimeMetricsConsumer:           uptimeMetricsConsumer,
			domainExpirationMetricsConsumer: domainExpirationMetricsConsumer,
		},
		nil
}

func (m *Manager) Run() {
	err := make(chan error)
	go m.sslTargetsManager(err)
	go m.uptimeTargetsManager(err)
	go m.domainExpirationTargetsManager(err)
	go m.sslMetricsConsumerManager(err)
	go m.uptimeMetricsConsumerManager(err)
	go m.domainExpirationMetricsConsumerManager(err)
	for {
		select {
		case <-m.svc.ctx.Done():
			return
		case e := <-err:
			if e != nil {
				log.Println(e)
			}
		}
	}
}
func (m *Manager) sslTargetsManager(err chan error) {
	for {
		select {
		case <-m.svc.ctx.Done():
			return
		case <-time.After(timePeriodSslTargets):
			targets, e := m.svc.TargetService.GetTargetsForSslChecker()
			if err != nil {
				err <- e
			}
			for _, target := range targets {
				msg, e := json.Marshal(target)
				if err != nil {
					err <- e
				}
				e = m.sslDataProducer.Publish(msg)
				if e != nil {
					err <- e
				}
			}
		}
	}
}

func (m *Manager) uptimeTargetsManager(err chan error) {
	for {
		select {
		case <-m.svc.ctx.Done():
			return
		case <-time.After(timePeriodUptimeTargets):
			targets, e := m.svc.TargetService.GetTargetsForUptimeChecker()
			if err != nil {
				err <- e
			}
			for _, target := range targets {
				msg, e := json.Marshal(target)
				if err != nil {
					err <- e
				}
				e = m.uptimeDataProducer.Publish(msg)
				if e != nil {
					err <- e
				}
			}
		}
	}
}

func (m *Manager) domainExpirationTargetsManager(err chan error) {
	for {
		select {
		case <-m.svc.ctx.Done():
			return
		case <-time.After(timePeriodDomainExpirationTargets):
			targets, e := m.svc.TargetService.GetTargetsForDomainExpirationChecker()
			if err != nil {
				err <- e
			}
			for _, target := range targets {
				msg, e := json.Marshal(target)
				if err != nil {
					err <- e
				}
				e = m.domainExpirationDataProducer.Publish(msg)
				if e != nil {
					err <- e
				}
			}
		}
	}
}

func (m *Manager) sslMetricsConsumerManager(err chan error) {
	messages := make(chan string)
	go m.sslMetricsConsumer.Consume(messages, err)
	for {
		select {
		case <-m.svc.ctx.Done():
			return
		case msg := <-messages:
			var d *entities.SslData
			e := json.Unmarshal([]byte(msg), &d)
			if e != nil {
				err <- e
				continue
			}
			e = m.svc.MetricsService.InsertSsl(d)
			if e != nil {
				err <- e
			}
		}
	}
}

func (m *Manager) uptimeMetricsConsumerManager(err chan error) {
	messages := make(chan string)
	go m.uptimeMetricsConsumer.Consume(messages, err)
	for {
		select {
		case <-m.svc.ctx.Done():
			return
		case msg := <-messages:
			var d *entities.UptimeData
			e := json.Unmarshal([]byte(msg), &d)
			if e != nil {
				err <- e
				continue
			}
			e = m.svc.MetricsService.InsertUptime(d)
			if e != nil {
				err <- e
			}
		}
	}
}

func (m *Manager) domainExpirationMetricsConsumerManager(err chan error) {
	messages := make(chan string)
	go m.domainExpirationMetricsConsumer.Consume(messages, err)
	for {
		select {
		case <-m.svc.ctx.Done():
			return
		case msg := <-messages:
			var d *entities.DomainExpirationData
			e := json.Unmarshal([]byte(msg), &d)
			if e != nil {
				err <- e
				continue
			}
			e = m.svc.MetricsService.InsertDomainExpiration(d)
			if e != nil {
				err <- e
			}
		}
	}
}

func (m *Manager) Stop() error {
	m.svc.ctx.Done()
	err := m.sslDataProducer.Close()
	if err != nil {
		return err
	}
	err = m.uptimeDataProducer.Close()
	if err != nil {
		return err
	}
	err = m.domainExpirationDataProducer.Close()
	if err != nil {
		return err
	}
	err = m.sslMetricsConsumer.Close()
	if err != nil {
		return err
	}
	err = m.uptimeMetricsConsumer.Close()
	if err != nil {
		return err
	}
	err = m.domainExpirationMetricsConsumer.Close()
	if err != nil {
		return err
	}
	return nil
}
