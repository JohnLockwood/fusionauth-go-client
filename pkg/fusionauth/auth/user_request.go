package auth

// UserRequest is the payload for a CreateUser request, for example.
// https://fusionauth.io/docs/v1/tech/apis/users#create-a-user
type UserRequest struct {
	SendSetPasswordEmail bool `json:"sendSetPasswordEmail"`
	SkipVerification     bool `json:"skipVerification"`
	User                 User `json:"user"`
}
