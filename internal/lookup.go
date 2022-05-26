package internal

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/natemarks/postgr8/credentials"
)


func GetTestCredentials()(creds credentials.CdkRdsAutoCredential, err error){

	// set secret list filter using the tags set in deployments/app.py
	testSecretInput := &secretsmanager.ListSecretsInput{
		Filters: []types.Filter{
			{
				Key:   "tag-key",
				Values: []string{"purpose"},
			},
			{
				Key:   "tag-value",
				Values: []string{"postgr8_test_fixture"},
			},
		},
	}
	// Get a lsit of secrets that match the filter
	listOutput, err := credentials.ListSecrets(testSecretInput)
	if err != nil {
		return creds, err
	}
	// Make sure there is exactly one match
	if len(listOutput) == 0 {
		return creds, errors.New(
			"no matching secrets. Check AWS credentials or deploy test fixture CDK")
	}
	if len(listOutput) > 1 {
		return creds, errors.New(
			"too many matching secrets. Clean up extras")
	}
	// Get the secrect name from the matching secret and use it to get the 
	// credentials
	secretID := *listOutput[0].Name
	creds, err = credentials.GetCredentialsFromSecretID(secretID)
	return creds, err
}