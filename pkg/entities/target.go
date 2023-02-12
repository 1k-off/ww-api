package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Target struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	URL              string             `bson:"url" json:"url" validate:"required,url"`
	IsActive         bool               `bson:"isActive" json:"isActive" validate:"required,boolean"`
	Uptime           bool               `bson:"uptime" json:"uptime" validate:"required,boolean"`
	SSL              bool               `bson:"ssl" json:"ssl" validate:"required,boolean"`
	DomainExpiration bool               `bson:"domainExpiration" json:"domainExpiration" validate:"required,boolean"`
	Config           TargetConfig       `bson:"config" json:"config"`
}

type TargetConfig struct {
	Uptime        UptimeConfig        `bson:"uptime" json:"uptime"`
	Notifications TargetNotifications `bson:"notifications" json:"notifications"`
}

type UptimeConfig struct {
	KeywordCheck bool   `bson:"keywordCheck" json:"keywordCheck"`
	KeywordType  string `bson:"keywordType" json:"keywordType"`
	Keyword      string `bson:"keyword" json:"keyword"`
}

type TargetNotifications struct {
	Slack      bool     `bson:"slack" json:"slack"`
	SlackUsers []string `bson:"slackUsers" json:"slackUsers"`
}

type SslTarget struct {
	URL string `json:"url,omitempty"`
}
