package secrets_manager

import (
	"ssp-portal-reporting-processor/config"
	"ssp-portal-reporting-processor/model"
	"ssp-portal-reporting-processor/utils"
)

func RetrieveAllDbData(dbConfig *config.DBConfig) (*model.DbDetails, error) {
	mobileAdSecretString, err := retrieveSecret(dbConfig.MobileAdDb.Secret.Region, dbConfig.MobileAdDb.Secret.Name)
	if err != nil {
		return nil, err
	}
	mobileAdDbInstances, err := utils.UnmarshalJson[model.DbInstances](mobileAdSecretString)
	if err != nil {
		return nil, err
	}
	tvAdSecretString, err := retrieveSecret(dbConfig.TvAdDb.Secret.Region, dbConfig.TvAdDb.Secret.Name)
	if err != nil {
		return nil, err
	}
	tvAdDbInstances, err := utils.UnmarshalJson[model.DbInstances](tvAdSecretString)
	if err != nil {
		return nil, err
	}
	return &model.DbDetails{
		MobileAd: *mobileAdDbInstances,
		TvAd:     *tvAdDbInstances,
	}, nil
}

func RetrieveEmailSecret(emailConfig *config.EmailConfig) (*model.SesCredentials, error) {
	emailSecretString, err := retrieveSecret(emailConfig.Secret.Region, emailConfig.Secret.Name)
	if err != nil {
		return nil, err
	}
	sesCredentials, err := utils.UnmarshalJson[model.SesCredentials](emailSecretString)
	if err != nil {
		return nil, err
	}
	sesCredentials.MailFrom = emailConfig.MailFrom
	return sesCredentials, nil
}
