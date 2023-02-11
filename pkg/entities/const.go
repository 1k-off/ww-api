package entities

const (
	MongoKeyId          = "_id"
	MongoKeyLogin       = "login"
	MongoKeyUrl         = "url"
	MongoKeyIsActive    = "isActive"
	MongoKeyMetadataUrl = "metadata.url"
	MongoKeyTimestamp   = "timestamp"
	MongoKeyMetadata    = "metadata"
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
