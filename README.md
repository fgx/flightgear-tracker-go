flightgear-tracker-go
============================================

-- Intro 

The fgtracker is currently sponsored by Hazuki in HK (thankyou hazuki)
- fgms is running on that machine and tracker enabled
- the tracker is periodically inserting rows into a postgres db
- the posttres db is on SSD discs (thankyou hazuki) and works
- the front end website is on a dark corer of the internet here.
--  http://mpserver15.flightgear.org/modules/fgtracker/


This is WIP and the initial working is the get an Ajax feed.

Goals
=======================
Shortterm - get access to realtime data ie this weeks stuff
- get this server running and quering against the live database for testing filtered big time
- create an ajax feed of data
- cache the data somehow so give pg a break
- make install and update easy and automated
- promote: live data refreshed every few seconds via http or a catch up


Midterm - consolidate data  and de-centralaise
- make the system http/websocket or UDP and itegrate crossfeed
- eliminate all crap data 
- make a way to archive legacy data, and serialise to kml, csv, git, or any format 
- allow VA filters
- all







Working
===================

- /flights/{callsign}
- /flight/{flight_id}


AJAX API
===================

Here is the suggessted api and is subject to change

 - /flights - returns flights dashboard
 - /flights/open - returns only open flights
 
 
 
 - /flight/{flight_id} - return the flight details and waypoints 
 
 
 - /callsign/{callsign} - returns callsign dashboard
 - /callsign/{callsign}/flights
 



