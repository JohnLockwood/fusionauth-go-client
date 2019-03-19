package auth

import (
	"testing"

	. "github.com/JohnLockwood/fusionauth-go-client/pkg/fusionauth/auth"
)

const (
	testAPIKey string = "6b87a398-39f2-4692-927b-13188a81a9a3"
	url        string = "http://localhost/9011"
)

func TestCreateClient(t *testing.T) {
	client := FusionAuthClient{URL: url, APIKey: testAPIKey}
	if client.APIKey != testAPIKey {
		t.Errorf("FusionAuthClient.APIKey = %s, want %s", client.APIKey, testAPIKey)
	}
}

func TestCreateUser(t *testing.T) {
	client := FusionAuthClient{URL: url, APIKey: testAPIKey}
	user := User{Email: "johndoe@example.com", Password: "Secret"}
	_, err := client.CreateUser(user, "")
	if err == nil {
		t.Errorf("TODO not implemented.  err should not be nil but for now it's expected.")
	}
}
