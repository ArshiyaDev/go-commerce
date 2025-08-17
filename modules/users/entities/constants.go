// Package users provides functionality for managing users in the system.
package entities

type Status string

const (
	Active   Status = "active"
	Inactive Status = "inactive"
	Verified Status = "verified"
	Suspend  Status = "suspend"
)
