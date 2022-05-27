package command_test

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/natemarks/postgr8/command"
	"github.com/natemarks/postgr8/internal"
)

func TestValidCredentials(t *testing.T) {
	connParams, err := internal.GetTestConnParams()
	if err != nil {
		t.Error(err)
	}
	type args struct {
		connData command.InstanceConnectionParams
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
		wantErr    bool
	}{
		{name: "valid",
			args:       args{connParams},
			wantResult: true,
			wantErr:    false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := command.ValidCredentials(tt.args.connData)
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
