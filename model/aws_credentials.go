package model

type DbConnection struct {
	Port     int    `json:"port"`
	Host     string `json:"host"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type DbInstances struct {
	PIData    DbConnection `json:"pi_data"`
	Secondary DbConnection `json:"secondary"`
	Primary   DbConnection `json:"primary"`
}

type DbDetails struct {
	TvAd     DbInstances
	MobileAd DbInstances
}

type SesCredentials struct {
	Arn      string `json:"arn"`
	MailFrom string `json:"-"`
}
