
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

type Flight struct {
	FlightId int64 `json:"flight_id"`
	Callsign string `json:"callsign"`
	Status string `json:"status"`
	Model string `json:"model"`
	Aero string `json:"aero"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	Duration string `json:"duration"`
	DurationSecs int64 `json:"duration_sec"`
}

type Waypoint struct {
	Time time.Time `json:"time"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	AltFt float64 `json:"alt_ft"`
	
}


type JSONFlightsPayload struct {
	Success bool `json:"success"`
	Flights []Flight `json:"flights"`
}
type JSONFlightPayload struct {
	Success bool `json:"success"`
	Flight Flight `json:"flight"`
	Track []Waypoint `json:"track"`
}



func GetFlights(callsign string)([]Flight, error){
	
	sql := " SELECT flights.id as flight_id, flights.callsign, flights.status, "
	sql += " flights.model, COALESCE(models.human_string, flights.model) as aero, "
	sql += " flights.start_time, flights.end_time, "
	sql += " to_char(end_time -  start_time, 'HH24:MI:SS') as duration_seconds,  "
	sql += " EXTRACT(EPOCH from end_time) - EXTRACT(EPOCH from start_time) as duration "
	sql += " from flights "
	sql += " left join models on flights.model = models.fg_string "
	sql += " WHERE flights.callsign =  $1"
	rows, query_err := Db.Query(sql, callsign)
	if query_err != nil {
		return nil, query_err
	}
	defer rows.Close()
	
	var flights []Flight
	for rows.Next() {
		var flight Flight			
		err_scan := rows.Scan(	&flight.FlightId, &flight.Callsign, &flight.Status, 
								&flight.Model, &flight.Aero, 
								&flight.StartTime, &flight.EndTime, &flight.Duration, &flight.DurationSecs)
		if err_scan != nil {
			return nil, err_scan
		}
		flights = append(flights, flight)
	}
	return flights, nil
}


// Return a flight from DB if found
// TODO: create view
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


func GetFlightTrack(flight_id int64)([]Waypoint, error){
	
	sql := " SELECT time, latitude, longitude, altitude " 
	sql += " FROM waypoints_all "
	sql += " WHERE flight_id=$1 " // AND (longitude!=0 OR latitude!=0 OR altitude!=0) 
	sql += " ORDER BY time "
	rows, query_err := Db.Query(sql,	flight_id)
	if query_err != nil {
		return nil, query_err
	}
	defer rows.Close()
	
	var wps []Waypoint
	for rows.Next() {
		var wp Waypoint			
		err_scan := rows.Scan(&wp.Time, &wp.Lat, &wp.Lon, &wp.AltFt)
		if err_scan != nil {
			return nil, err_scan
		}
		wps = append(wps, wp)
	}
	return wps, nil
}

//= /flight/{flight_id}
func Flight_AjaxHandler(w http.ResponseWriter, r *http.Request){
	
	// check flight_id is an int type
	vars := mux.Vars(r)
	flight_id, id_error := strconv.ParseInt(vars["flight_id"], 10, 64)
	if id_error != nil {
		fmt.Fprint(w, CreateAjaxErrorPayload("Invalid flight_id:", nil))
		return
	}
	
	flight, flight_err := GetFlight(flight_id)
	if flight_err != nil {
		fmt.Fprint(w, CreateAjaxErrorPayload("DB Error:", flight_err))
		return
	}
	track, track_err := GetFlightTrack(flight_id)
	if track_err != nil {
		fmt.Fprint(w, CreateAjaxErrorPayload("DB Error:", track_err))
		return
	}
	
	payload := JSONFlightPayload{Success: true, Flight: flight, Track: track}
	//log.Println("flight", flight)
	s, err := json.Marshal(payload)
	if err != nil {
		log.Println("flight", err)
	}
	fmt.Fprint(w, string(s))
}


//= /flights/{callsign}
func FlightsByCallsign_AjaxHandler(w http.ResponseWriter, r *http.Request){
	
	// TODO check callsign is valid 
	vars := mux.Vars(r)
	callsign := vars["callsign"]
	
	flights, flights_err := GetFlights(callsign)
	if flights_err != nil {
		fmt.Fprint(w, CreateAjaxErrorPayload("DB Error:", flights_err))
		return
	}
	
	payload := JSONFlightsPayload{Success: true, Flights: flights}

	s, err := json.Marshal(payload)
	if err != nil {
		log.Println("flight", err)
	}
	fmt.Fprint(w, string(s))
}
