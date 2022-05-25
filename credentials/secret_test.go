package credentials

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
)

func TestSecrets(t *testing.T) {
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
	type args struct {
		input *secretsmanager.ListSecretsInput
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
	}{
		{name: "valid",
	args: args{
		input: testSecretInput,
	},wantErr: false,},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ListSecrets(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListSecrets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
