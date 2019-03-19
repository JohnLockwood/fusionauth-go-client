package auth

import (
	"net/url"
	"time"
)

// User of the system that can be created or authenticated
// For optional fields, semantics, etc. see
// https://fusionauth.io/docs/v1/tech/apis/users#create-a-user
// Sent as part of a UserRequest, see user_request.go
type User struct {
	Active                 bool              `json:"active"`
	BirthDate              time.Time         `json:"birthDate"`
	CleanSpeakID           string            `json:"cleanSpeakId"`
	Data                   map[string]string `json:"data"`
	Email                  string            `json:"email"`
	EncryptionScheme       string            `json:"encryptionScheme"`
	Expiry                 time.Time         `json:"expiry"`
	Factor                 string            `json:"factor"`
	FirstName              string            `json:"firstName"`
	FullName               string            `json:"fullName"`
	ImageURL               url.URL           `json:"imageUrl"`
	InsertInstant          time.Time         `json:"insertInstant"`
	LastLoginInstant       time.Time         `json:"lastLoginInstant"`
	LastName               string            `json:"lastName"`
	MiddleName             string            `json:"middleName"`
	MobilePhone            string            `json:"mobilePhone"`
	Password               string            `json:"password"`
	PasswordChangeRequired bool              `json:"passwordChangeRequired"`
	PreferredLanguages     []string          `json:"preferredLanguages"`
	Registrations          []Registration    `json:"registrations"`
	TenantID               string            `json:"tenantId"`
	Timezone               string            `json:"timezone"`
	TwoFactorDelivery      TwoFactorDelivery `json:"twoFactorDelivery"`
	TwoFactorEnabled       bool              `json:"twoFactorEnabled"`
	TwoFactorSecret        string            `json:"twoFactorSecret"`
	Username               string            `json:"username"`
	UsernameStatus         string            `json:"usernameStatus"`
}
