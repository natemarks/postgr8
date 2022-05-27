package sql

import (
	"fmt"
)

// TableRowCount slow count of rows TRASH
// SELECT count(*) FROM tableName;
func TableRowCount(tableName string) string {
	return fmt.Sprintf("SELECT count(*) FROM %s;", tableName)
}
