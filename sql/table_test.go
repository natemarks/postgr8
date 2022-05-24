package sql_test

import (
	"testing"

	"github.com/natemarks/postgr8/sql"
)

// TestTableRowCount TRASH
func TestTableRowCount(t *testing.T) {
	type args struct {
		tableName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid",
			args: args{
				tableName: "asd123",
			},
			want: "SELECT count(*) FROM asd123;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sql.TableRowCount(tt.args.tableName); got != tt.want {
				t.Errorf("TableRowCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
