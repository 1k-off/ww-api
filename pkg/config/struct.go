package config

type Config struct {
	LogLevel string   `mapstructure:"log_level"`
	Database Database `mapstructure:"database"`
	Server   Server   `mapstructure:"server"`
	Queue    Queue    `mapstructure:"queue"`
}

type Database struct {
	Provider         string `mapstructure:"provider"`
	ConnectionString string `mapstructure:"connection_string"`
}
type Server struct {
	Port         string   `mapstructure:"port"`
	AccessToken  ServerAT `mapstructure:"access_token"`
	RefreshToken ServerRt `mapstructure:"refresh_token"`
}

type Queue struct {
	Provider string       `mapstructure:"provider"`
	Memphis  QueueMemphis `mapstructure:"memphis"`
}

type ServerAT struct {
	PrivateKey string `mapstructure:"private_key"`
	PublicKey  string `mapstructure:"public_key"`
	ExpiresIn  int    `mapstructure:"expires_in"`
}
type ServerRt struct {
	PrivateKey string `mapstructure:"private_key"`
	PublicKey  string `mapstructure:"public_key"`
	ExpiresIn  int    `mapstructure:"expires_in"`
}

type QueueMemphis struct {
	ClientLogin               string `mapstructure:"client_login"`
	ClientToken               string `mapstructure:"client_token"`
	Url                       string `mapstructure:"url"`
	SslTargetsSN              string `mapstructure:"ssl_targets_station_name"`
	UptimeTargetsSN           string `mapstructure:"uptime_targets_station_name"`
	DomainExpirationTargetsSN string `mapstructure:"domain_expiration_targets_station_name"`
	SslMetricsSN              string `mapstructure:"ssl_metrics_station_name"`
	UptimeMetricsSN           string `mapstructure:"uptime_metrics_station_name"`
	DomainExpirationMetricsSN string `mapstructure:"domain_expiration_metrics_station_name"`
}
