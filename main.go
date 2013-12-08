
package main

import (
    "log"
    "fmt"    
    _ "github.com/lib/pq"
	"database/sql"
	
    "github.com/fgx/flightgear-tracker-go/config"
    "github.com/fgx/flightgear-tracker-go/tracker"
)

func main(){

	//= Load config (LoadConfig will sysexit if invalid)
	conf := config.LoadConfig()
	log.Println("open config: ", conf.Database)

	//= Create DB Connection (navdata.Db is the connection pointer
	url := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", conf.DbUser, conf.DbPassword, conf.DbServer, conf.Database)
	log.Println( url )
	
	var err error
	tracker.Db, err = sql.Open("postgres", "postgres://mash2:mash2@localhost/fgxmap?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer tracker.Db.Close()
}