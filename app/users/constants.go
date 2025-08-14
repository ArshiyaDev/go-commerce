package users

type Status string

const (
	Active   Status = "active"
	Inactive Status = "inactive"
	Verified Status = "verified"
	Suspend  Status = "suspend"
)
