package auth

// RegistrationRequest creates a registration for a new or existing User
type RegistrationRequest struct {
	GenerateAuthenticationToken  bool         `json:"generateAuthenticationToken"`
	Registration                 Registration `json:"registration"`
	SendSetPasswordEmail         bool         `json:"sendSetPasswordEmail"`
	SkipRegistrationVerification bool         `json:"skipRegistrationVerification"`
	SkipVerification             bool         `json:"skipVerification"`
	User                         User         `json:"user"`
}
