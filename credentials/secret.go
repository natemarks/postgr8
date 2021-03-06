package credentials

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/natemarks/postgr8/command"
)

// ListSecrets Use ListSecretsInput to get a slice of secret entries
// This really just handles the pagination for me
func ListSecrets(input *secretsmanager.ListSecretsInput) (secretList []types.SecretListEntry, err error) {
	//var input secretsmanager.ListSecretsInput
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return secretList, err
	}

	client := *secretsmanager.NewFromConfig(cfg)

	paginator := *secretsmanager.NewListSecretsPaginator(
		&client,
		input,
	)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			return secretList, err
		}
		secretList = append(secretList, output.SecretList...)
	}

	return secretList, err
}

// GetCredentialsFromSecretID Given secretID return command.InstanceConnectionParams
func GetCredentialsFromSecretID(secretID string) (credentials command.InstanceConnectionParams, err error) {
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
