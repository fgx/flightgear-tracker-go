DROP VIEW v_flights;

CREATE OR REPLACE VIEW v_flights AS
SELECT flights.id AS flight_id, 
flights.callsign, flights.status, 
flights.model, COALESCE(models.human_string, flights.model) AS aero, 
flights.start_time, flights.end_time, 
to_char(end_time -  start_time, 'HH24:MI:SS') AS duration,  
EXTRACT(EPOCH FROM end_time) - EXTRACT(EPOCH FROM start_time) AS duration_sec 
FROM flights
LEFT JOIN models ON flights.model = models.fg_string 
;