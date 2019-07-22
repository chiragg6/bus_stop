package bus

import (
	"errors"
)

var directionMap map[string]string

const (
	routeURL = "http://svc.metrotransit.org/NexTrip/Routes"
	//http://svc.metrotransit.org/NexTrip/Directions/{ROUTE}
	directionURL = "http://svc.metrotransit.org/NexTrip/Directions/"
	//http://svc.metrotransit.org/NexTrip/Stops/{ROUTE}/{DIRECTION}
	stopsURL = "http://svc.metrotransit.org/NexTrip/Stops/"
	//http://svc.metrotransit.org/NexTrip/{ROUTE}/{DIRECTION}/{STOP}
	timeDeparturesURL = "http://svc.metrotransit.org/NexTrip/"
)

var noBusError error

func init() {
	directionMap = make(map[string]string)
	//1 = South, 2 = East, 3 = West, 4 = North.
	directionMap["south"] = "1"
	directionMap["east"] = "2"
	directionMap["west"] = "3"
	directionMap["north"] = "4"

	noBusError = errors.New("no bus found")
}
