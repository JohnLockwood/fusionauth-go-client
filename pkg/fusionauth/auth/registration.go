package auth

// Registration links a User to an Application
// TODO this object is not complete
type Registration struct {
	ApplicationID      string            `json:"applicationId"`
	Data               map[string]string `json:"data"`
	PreferredLanguages []string          `json:"preferredLanguages"`
	ID                 string            `json:"id"`
	Roles              []string          `json:"roles"`
	Timezone           string            `json:"timezone"`
	Username           string            `json:"username"`
}
