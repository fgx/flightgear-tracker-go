
package main

import (
    "log"
    "fmt"   
     
    _ "github.com/lib/pq"
	"database/sql"
	
	"net/http"
	"github.com/codegangsta/martini"
	
    "github.com/fgx/flightgear-tracker-go/config"
    "github.com/fgx/flightgear-tracker-go/tracker"
)

func main(){

	//= Load config (LoadConfig will sys exit(1) if error)
	conf := config.LoadConfig()
	log.Println("Loaded config: ", conf)

	//= Create DB Connection (tracker.Db is the connection pointer
	url := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", conf.DbUser, conf.DbPassword, conf.DbServer, conf.Database)	
	var err error
	tracker.Db, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	//= Test db connection
	con_err := tracker.Db.Ping()
	if con_err != nil {
		log.Fatal("DB error: (check config)", con_err)
	}
	log.Println("DB connected: OK ")
	defer tracker.Db.Close()
	
	//= Setup Martini and routing
	m := martini.Classic()
	m.Get("/", func() string {
		return "yes"
	})
	
	// Lets go !
	log.Println("Listening on: ", conf.HttpPort)
	http.ListenAndServe(fmt.Sprintf(":%d", conf.HttpPort) , m)
	
	
}