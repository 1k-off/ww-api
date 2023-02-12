package config

import (
	"errors"
	"github.com/spf13/viper"
)

const (
	DbProviderMongodb = "mongodb"
)

const (
	QueueProviderMemphis = "memphis"
)

const (
	AccessTokenExpiresIn  = 900
	RefreshTokenExpiresIn = 3600

	MemphisSslTargetsStationName              = "ssl-targets"
	MemphisUptimeTargetsStationName           = "uptime-targets"
	MemphisDomainExpirationTargetsStationName = "domain-expiration-targets"
	MemphisSslMetricsStationName              = "ssl-metrics"
	MemphisUptimeMetricsStationName           = "uptime-metrics"
	MemphisDomainExpirationMetricsStationName = "domain-expiration-metrics"
)

func newDefaultConfig() *Config {
	return &Config{
		Database: Database{
			Provider: DbProviderMongodb,
		},
		Server: Server{
			Port: "8080",
			AccessToken: ServerAT{
				ExpiresIn: AccessTokenExpiresIn,
			},
			RefreshToken: ServerRt{
				ExpiresIn: RefreshTokenExpiresIn,
			},
		},
		Queue: Queue{
			Provider: QueueProviderMemphis,
			Memphis: QueueMemphis{
				SslTargetsSN:              MemphisSslTargetsStationName,
				UptimeTargetsSN:           MemphisUptimeTargetsStationName,
				DomainExpirationTargetsSN: MemphisDomainExpirationTargetsStationName,
				SslMetricsSN:              MemphisSslMetricsStationName,
				UptimeMetricsSN:           MemphisUptimeMetricsStationName,
				DomainExpirationMetricsSN: MemphisDomainExpirationMetricsStationName,
			},
		},
	}
}

func Load(path string) (*Config, error) {
	cfg := newDefaultConfig()

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) Validate() error {
	if err := c.validateDbConfig(); err != nil {
		return err
	}
	if err := c.validateQueueConfig(); err != nil {
		return err
	}
	return nil
}

func (c *Config) validateDbConfig() error {
	switch c.Database.Provider {
	case DbProviderMongodb:
		if c.Database.ConnectionString == "" {
			return errors.New("database configuration invalid")
		}
	default:
		return errors.New("database provider is not supported or invalid")
	}
	return nil
}

func (c *Config) validateQueueConfig() error {
	switch c.Queue.Provider {
	case QueueProviderMemphis:
		if c.Queue.Memphis.ClientLogin == "" || c.Queue.Memphis.ClientToken == "" || c.Queue.Memphis.Url == "" {
			return errors.New("queue configuration invalid")
		}
	default:
		return errors.New("queue provider is not supported or invalid")
	}
	return nil
}
