package entities

const (
	MongoKeyId        = "_id"
	MongoKeyLogin     = "login"
	MongoKeyUrl       = "url"
	MongoKeyIsActive  = "isActive"
	MongoKeyTimestamp = "timestamp"
	MongoKeyMetadata  = "metadata"

	MongoKeyUptime           = "uptime"
	MongoKeySsl              = "ssl"
	MongoKeyDomainExpiration = "domainExpiration"

	MongoKeyMetricMetadataUrl  = "metadata.url"
	MongoKeyMetricUptimeUp     = "up"
	MongoKeyMetricExpiringSoon = "expiringSoon"
	MongoKeyMetricError        = "error"
)
const (
	MongoTsGranularitySeconds = "seconds"
	MongoTsGranularityHours   = "hours"
)
const (
	MongoCollectionNameUsers                   = "users"
	MongoCollectionNameTargets                 = "targets"
	MongoCollectionNameMetricsSsl              = "metrics_ssl"
	MongoCollectionNameMetricsUptime           = "metrics_uptime"
	MongoCollectionNameMetricsDomainExpiration = "metrics_domain_expiration"
)
const (
	CheckerNameSsl              = "ssl"
	CheckerNameUptime           = "uptime"
	CheckerNameDomainExpiration = "domainExpiration"
)
