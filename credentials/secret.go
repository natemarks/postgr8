package credentials

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// CdkRdsAutoCredential When CDK deploys an RDS instances and automatically
// generates  credentials in secretsmanager, this is the format of the JSON
type CdkRdsAutoCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Engine   string `json:"engine"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

// GetCredentialsFromSecretsmanager Given secretID return CdkRdsAutoCredential
func GetCredentialsFromSecretsmanager(secretID string) (credentials CdkRdsAutoCredential, err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return credentials, err
	}

	client := *secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretID),
	}
	result, err := client.GetSecretValue(context.TODO(), input)
	if err != nil {
		return credentials, err
	}

	err = json.Unmarshal([]byte(*result.SecretString), &credentials)
	return credentials, err
}
