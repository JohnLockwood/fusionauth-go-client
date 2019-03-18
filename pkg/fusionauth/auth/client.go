package auth

import "fmt"

type FusionAuthClient struct {
	URL    string
	APIKey string
}

type Request interface{}

func (*FusionAuthClient) CreateUser(request Request) (ClientResponse, error) {
	return nil, fmt.Errorf("Not implemented")
}

// Rest client

type ClientResponse interface{}
