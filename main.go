
package main

import (
    "log"
    "fmt"   
     
    _ "github.com/lib/pq"
	"database/sql"
	
	"net/http"
	"github.com/gorilla/mux"
	
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
	

		
	//= Setup routing
	router := mux.NewRouter()
	//m.Get("/", func() string {
	//	return "yes"
	//})
	router.HandleFunc("/flights/{callsign}", tracker.FlightsByCallsign_AjaxHandler)
	router.HandleFunc("/flight/{flight_id}", tracker.Flight_AjaxHandler)
	
	http.Handle("/", router)
	
	//= Start Server
	server_address := fmt.Sprintf(":%d", conf.HttpPort)
	log.Println("Listening on: ", server_address)
    if err := http.ListenAndServe(server_address, nil); err != nil {
		panic(err)
	}
	
	// Lets go !
	log.Println("Listening on: ", conf.HttpPort)
	
	
}