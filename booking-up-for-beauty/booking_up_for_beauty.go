// Package booking provides functions for booking appointments and managing dates.
package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	appointment, _ := time.Parse(layout, date)
	return appointment
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	appointment, _ := time.Parse(layout, date)
	return appointment.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	appointment, _ := time.Parse(layout, date)
	appointmentHour := appointment.Hour()
	return appointmentHour >= 12 && appointmentHour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	appointment, _ := time.Parse(layout, date)
	return fmt.Sprintf("You have an appointment on %s.", appointment.Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	layout := "2006-01-02 15:04:05 -0700 MST"
	date := fmt.Sprintf("%d-09-15 00:00:00 +0000 UTC", time.Now().Year())
	anniversary, _ := time.Parse(layout, date)
	return anniversary
}
