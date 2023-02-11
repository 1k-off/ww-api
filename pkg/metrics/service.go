package metrics

import "ww-api/pkg/entities"

type Service interface {
	InsertUptime(d *entities.UptimeData) error
	InsertUptimeBatch(d []*entities.UptimeData) error
	InsertSsl(d *entities.SslData) error
	InsertSslBatch(d []*entities.SslData) error
	InsertDomainExpiration(d *entities.DomainExpirationData) error
	InsertDomainExpirationBatch(d []*entities.DomainExpirationData) error
	Delete(url string)
	GetDownTargets() ([]*entities.TargetDown, error)
	GetSslExpiringSoon() ([]*entities.SslExpiringSoon, error)
	GetDomainExpiringSoon() ([]*entities.DomainExpiringSoon, error)
	GetStats() (*entities.MetricsStats, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertUptime(d *entities.UptimeData) error {
	return s.repository.InsertUptime(d)
}
func (s *service) InsertUptimeBatch(d []*entities.UptimeData) error {
	dI := make([]interface{}, len(d))
	for k, v := range d {
		dI[k] = v
	}
	return s.repository.InsertUptime(dI)
}

func (s *service) InsertSsl(d *entities.SslData) error {
	return s.repository.InsertSsl(d)
}
func (s *service) InsertSslBatch(d []*entities.SslData) error {
	dI := make([]interface{}, len(d))
	for k, v := range d {
		dI[k] = v
	}
	return s.repository.InsertSsl(dI)
}

func (s *service) InsertDomainExpiration(d *entities.DomainExpirationData) error {
	return s.repository.InsertDomainExpiration(d)
}
func (s *service) InsertDomainExpirationBatch(d []*entities.DomainExpirationData) error {
	dI := make([]interface{}, len(d))
	for k, v := range d {
		dI[k] = v
	}
	return s.repository.InsertDomainExpiration(dI)
}

func (s *service) Delete(url string) {
	s.repository.Delete(url)
}

func (s *service) GetDownTargets() ([]*entities.TargetDown, error) {
	return s.repository.GetDownTargets()
}

func (s *service) GetSslExpiringSoon() ([]*entities.SslExpiringSoon, error) {
	return s.repository.GetSslExpiringSoon()
}

func (s *service) GetDomainExpiringSoon() ([]*entities.DomainExpiringSoon, error) {
	return s.repository.GetDomainExpiringSoon()
}

func (s *service) GetStats() (*entities.MetricsStats, error) {
	return s.repository.GetStats()
}
