package app

import (
	"encoding/json"
	"log"
	"time"
	"ww-api/pkg/queue"
	"ww-api/pkg/queue/memphis"
)

type Manager struct {
	svc                          *Service
	sslDataProducer              queue.Producer
	uptimeDataProducer           queue.Producer
	domainExpirationDataProducer queue.Producer
}

func (s *Service) NewManager(mUser, mToken, mUrl, producerName, sslDSN, uptimeDSN, domainExpirationDSN string) (*Manager, error) {
	sslDataProducer, err := memphis.NewProducer(mUser, mToken, mUrl, sslDSN, producerName)
	if err != nil {
		return nil, err
	}
	uptimeDataProducer, err := memphis.NewProducer(mUser, mToken, mUrl, uptimeDSN, producerName)
	if err != nil {
		return nil, err
	}
	domainExpirationDataProducer, err := memphis.NewProducer(mUser, mToken, mUrl, domainExpirationDSN, producerName)
	if err != nil {
		return nil, err
	}
	return &Manager{
			svc:                          s,
			sslDataProducer:              sslDataProducer,
			uptimeDataProducer:           uptimeDataProducer,
			domainExpirationDataProducer: domainExpirationDataProducer,
		},
		nil
}

func (m *Manager) Run() {
	err := make(chan error)
	go m.sslTargetsManager(err)
	go m.uptimeTargetsManager(err)
	go m.domainExpirationTargetsManager(err)
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
		case <-time.After(time.Minute * 1):
			targets, e := m.svc.TargetService.GetTargetsForChecker("ssl")
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
		case <-time.After(time.Minute * 1):
			targets, e := m.svc.TargetService.GetTargetsForChecker("uptime")
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
		case <-time.After(time.Minute * 1):
			targets, e := m.svc.TargetService.GetTargetsForChecker("domainExpiration")
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
