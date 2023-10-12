package config

import (
	"slices"

	"github.com/spf13/viper"
)

type Config struct {
	DBConfig    DBConfig    `mapstructure:"db"`
	S3Config    S3Config    `mapstructure:"s3"`
	EmailConfig EmailConfig `mapstructure:"email"`
}

type Secret struct {
	Name   string `mapstructure:"name"`
	Region string `mapstructure:"region"`
}

type DBConfig struct {
	// Database configuration fields
	TvAdDb     DbSecret `mapstructure:"tvad"`
	MobileAdDb DbSecret `mapstructure:"mobilead"`
}

type DbSecret struct {
	Secret Secret `mapstructure:"secret"`
}

type S3Config struct {
	// S3 configuration fields
	Bucket                  string `mapstructure:"bucketName"`
	Region                  string `mapstructure:"region"`
	PresignedURLExpiryHours int    `mapstructure:"presignedUrlExpiryHours"`
}

type EmailConfig struct {
	// Email configuration fields
	MailFrom string `mapstructure:"mailFrom"`
	Secret   Secret `mapstructure:"secret"`
}

func LoadConfig(profile string) (*Config, error) {
	var cfg Config
	allProfiles := []string{"local", "dev", "stg", "stg-int", "prd"}
	isProfileCorrect := slices.Contains(allProfiles, profile)
	if !isProfileCorrect {
		return nil, viper.UnsupportedConfigError(profile)
	}
	viper.SetConfigName(profile)
	viper.AddConfigPath("config/resource")
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
