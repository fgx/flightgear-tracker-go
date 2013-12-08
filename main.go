
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

	//= Load config (LoadConfig will sys exit(1) if error)
	conf := config.LoadConfig()
	log.Println("opened config: ", conf.Database)

	//= Create DB Connection (tracker.Db is the connection pointer
	url := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", conf.DbUser, conf.DbPassword, conf.DbServer, conf.Database)	
	var err error
	tracker.Db, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	defer tracker.Db.Close()
}