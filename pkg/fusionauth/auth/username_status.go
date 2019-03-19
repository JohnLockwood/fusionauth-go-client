package auth

// UsernameStatus is used if moderating usernames using CleanSpeak
type UsernameStatus string

// Valid Values for UserNameStatus
const (
	ACTIVE   string = "ACTIVE"
	PENDING  string = "PENDING"
	REJECTED string = "REJECTED"
)
