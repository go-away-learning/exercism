// Package speed provides types and functions related to remote-controlled cars and tracks.
package speed

// Car represents a remote-controlled car.
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
	}
}

// Track represents a track for the remote-controlled car.
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	newBattery := car.battery - car.batteryDrain
	if newBattery < 0 {
		return car
	}
	car.battery = newBattery
	car.distance += car.speed

	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	distance := track.distance - car.distance
	batteryNeeded := (distance / car.speed) * car.batteryDrain
	return car.battery >= batteryNeeded
}
