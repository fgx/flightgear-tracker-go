
package tracker

import(
	_ "github.com/lib/pq"
	"database/sql"
)

//= Global Pointer initialised in main.go
var Db *sql.DB