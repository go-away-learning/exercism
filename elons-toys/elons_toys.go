// Package elon provides functionality for managing and displaying the state of a car.
package elon

import "fmt"

// Drive decreases the car's battery based on its battery drain rate and increases its distance traveled.
// If the battery would become negative, the method does nothing to avoid invalid battery levels.
func (c *Car) Drive() {
	newBattery := c.battery - c.batteryDrain
	if newBattery < 0 {
		return
	}
	c.battery = newBattery
	c.distance += c.speed
}

// DisplayDistance returns a string indicating the total distance traveled by the car.
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery returns a string indicating the current battery level of the car in percentage.
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish determines if the car can finish a given track distance with its current battery level.
func (c Car) CanFinish(trackDistance int) bool {
	distance := trackDistance - c.distance
	batteryNeeded := (float64(distance) / float64(c.speed)) * float64(c.batteryDrain)
	return float64(c.battery) >= batteryNeeded
}
