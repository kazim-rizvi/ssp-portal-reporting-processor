package secrets_manager

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func retrieveSecret(region string, secretName string) (string, error) {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		fmt.Println("Error creating AWS session:", err)
		return "", err
	}

	// Create a Secrets Manager client
	svc := secretsmanager.New(sess)

	// Get the secret value from Secrets Manager
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		fmt.Println("Error getting secret value:", err)
		return "", err
	}
	return *result.SecretString, nil

}
