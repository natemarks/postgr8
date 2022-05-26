package pg

import (
	"testing"

	"github.com/natemarks/postgr8/credentials"
	"github.com/natemarks/postgr8/internal"
)

func TestValidCredentials(t *testing.T) {
	fixtureCreds, err := internal.GetTestCredentials()
	if err != nil {
		t.Error(err)
	}
	type args struct {
		creds credentials.CdkRdsAutoCredential
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
		wantErr    bool
	}{
		{name: "valid",
			args:       args{fixtureCreds},
			wantResult: true,
			wantErr:    false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := ValidCredentials(tt.args.creds)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("ValidCredentials() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
