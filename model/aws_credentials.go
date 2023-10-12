package model

type DbConnection struct {
	Port     int    `json:"port"`
	Host     string `json:"host"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type AllDbConnections struct {
	PIData    DbConnection `json:"pi_data"`
	Secondary DbConnection `json:"secondary"`
	Primary   DbConnection `json:"primary"`
}

type SesCredentials struct {
	Arn string
}
