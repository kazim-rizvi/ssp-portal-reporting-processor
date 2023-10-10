package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBConfig    DBConfig    `mapstructure:"db"`
	CSVConfig   CSVConfig   `mapstructure:"csv"`
	S3Config    S3Config    `mapstructure:"s3"`
	EmailConfig EmailConfig `mapstructure:"email"`
}

type DBConfig struct {
	// Database configuration fields
	Host     string `mapstructure:"db_host"`
	Port     int    `mapstructure:"db_port"`
	Username string `mapstructure:"db_username"`
	Password string `mapstructure:"db_password"`
	Database string `mapstructure:"db_database"`
}

type CSVConfig struct {
	// CSV configuration fields
	OutputFilePath string
}

type S3Config struct {
	// S3 configuration fields
	AccessKey               string `mapstructure:"s3_access_key"`
	SecretKey               string `mapstructure:"s3_secret_key"`
	Bucket                  string `mapstructure:"s3_bucket"`
	Region                  string `mapstructure:"s3_region"`
	PresignedURLExpiryHours int    `mapstructure:"presigned_url_expiry_hours"`
}

type EmailConfig struct {
	// Email configuration fields
	SMTPServer      string   `mapstructure:"smtp_server"`
	SMTPPort        int      `mapstructure:"smtp_port"`
	SMTPUsername    string   `mapstructure:"smtp_username"`
	SMTPPassword    string   `mapstructure:"smtp_password"`
	SenderEmail     string   `mapstructure:"sender_email"`
	RecipientEmails []string `mapstructure:"recipient_emails"`
}

func LoadConfig(env string) (*Config, error) {
	var cfg Config
	// check for possible configs and throw error
	viper.SetConfigName(env)
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
