package entities

import "time"

type Metadata struct {
	Location string `json:"location" bson:"location"`
	URL      string `json:"url" bson:"url"`
}

type UptimeData struct {
	Metadata   Metadata  `json:"metadata" bson:"metadata"`
	Timestamp  time.Time `json:"timestamp" bson:"timestamp"`
	Up         bool      `json:"up" bson:"up"`
	StatusCode int       `json:"statusCode" bson:"statusCode"`
	Error      string    `json:"error" bson:"error"`
	IsChanged  bool      `json:"isChanged"`
}

type SslData struct {
	Metadata       Metadata    `json:"metadata" bson:"metadata"`
	Timestamp      time.Time   `json:"timestamp" bson:"timestamp"`
	ExpirationDate string      `json:"expirationDate" bson:"expirationDate"`
	CertData       SslCertData `json:"certData" bson:"certData"`
	Error          string      `json:"error" bson:"error"`
	ExpiringSoon   bool        `json:"expiringSoon" bson:"expiringSoon"`
}
type SslCertData struct {
	Host             string   `json:"host" bson:"host"`
	CommonName       string   `json:"commonName" bson:"commonName"`
	AlternativeNames []string `json:"alternativeNames" bson:"alternativeNames"`
	Issuer           string   `json:"issuer" bson:"issuer"`
	ValidFrom        string   `json:"validFrom" bson:"validFrom"`
	ValidTo          string   `json:"validTo" bson:"validTo"`
}

type DomainExpirationData struct {
	Metadata       Metadata  `json:"metadata" bson:"metadata"`
	Timestamp      time.Time `json:"timestamp" bson:"timestamp"`
	ExpirationDate string    `json:"expirationDate" bson:"expirationDate"`
	Error          string    `json:"error" bson:"error"`
	ExpiringSoon   bool      `json:"expiringSoon" bson:"expiringSoon"`
}

type TargetDown struct {
	Url   string    `json:"url" bson:"url"`
	Since time.Time `json:"since" bson:"since"`
}

type SslExpiringSoon struct {
	Url     string    `json:"url" bson:"url"`
	Expires time.Time `json:"since" bson:"since"`
}

type DomainExpiringSoon struct {
	Url     string    `json:"url" bson:"url"`
	Expires time.Time `json:"since" bson:"since"`
}

type MetricsStats struct {
	DownNow            int `json:"downNow" bson:"downNow"`
	SslExpiringSoon    int `json:"sslExpiringSoon" bson:"sslExpiringSoon"`
	DomainExpiringSoon int `json:"domainExpiringSoon" bson:"domainExpiringSoon"`
}
