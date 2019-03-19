package auth

import "fmt"

// FusionAuthClient represents a connection to FusionAuth
type FusionAuthClient struct {
	URL    string
	APIKey string
}

// CreateUser creates a user with an optional id.  If the id is an empty string,
// the system will create an ID for the user.
func (*FusionAuthClient) CreateUser(user User, id string) (ClientResponse, error) {
	return nil, fmt.Errorf("Not implemented")
}

// Rest client

// ClientResponse represents a successful response from an API call
type ClientResponse interface{}
