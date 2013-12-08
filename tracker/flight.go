
package tracker

import (
	"fmt"
	"log"
    "strconv"
    "time"
	//"database/sql"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type JSONFlightPayload struct {
	Success bool `json:"success"`
	Flight Flight
}

type Flight struct {
	FlightId int64 `json:"flight_id"`
	Callsign string `json:"callsign"`
	Status string `json:"status"`
	Model string `json:"model"`
	Aero string `json:"aero"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	Duration string `json:"duration"`
	DurationSecs int64 `json:"duration_secs"`
}


// Return a flight from DB if found
func GetFlight(flight_id int64)(Flight, error){
	var flight Flight
	sql := " SELECT flights.id as flight_id, flights.callsign, flights.status, "
	sql += " flights.model, COALESCE(models.human_string, flights.model) as aero, "
	sql += " flights.start_time, flights.end_time, "
	sql += " to_char(end_time -  start_time, 'HH24:MI:SS') as duration_seconds,  "
	sql += " EXTRACT(EPOCH from end_time) - EXTRACT(EPOCH from start_time) as duration "
	sql += " from flights "
	sql += " left join models on flights.model = models.fg_string "
	sql += " WHERE id =  $1"
	err := Db.QueryRow(sql,	flight_id).Scan(
						&flight.FlightId, &flight.Callsign, &flight.Status, 
						&flight.Model, &flight.Aero, 
						&flight.StartTime, &flight.EndTime, &flight.Duration, &flight.DurationSecs)
	return flight, err

}


func FlightHandler(w http.ResponseWriter, r *http.Request){
	
	// check flight_id is an int type
	vars := mux.Vars(r)
	flight_id, id_error := strconv.ParseInt(vars["flight_id"], 10, 64)
	if id_error != nil {
		fmt.Fprint(w, CreateAjaxErrorPayload("Invalid flight_id:", nil))
		return
	}
	flight, err := GetFlight(flight_id)
	if err != nil {
		fmt.Fprint(w, CreateAjaxErrorPayload("DB Error:", err))
		return// 200, "Some error"
	}
	payload := JSONFlightPayload{Success: true, Flight: flight}
	//log.Println("flight", flight)
	s, err := json.Marshal(payload)
	if err != nil {
		log.Println("flight", err)
	}
	//return 200, string(s)
	fmt.Fprint(w, string(s))
}

