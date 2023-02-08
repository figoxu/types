package types

import "time"

var timeZone = time.FixedZone("UTC+8", int((8 * time.Hour).Seconds()))

func SetZone(location *time.Location) {
	timeZone = location
}

func GetZone() *time.Location {
	return timeZone
}
