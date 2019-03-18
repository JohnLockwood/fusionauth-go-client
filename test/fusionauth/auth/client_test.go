package auth

import (
	"testing"

	. "github.com/JohnLockwood/fusionauth-go-client/pkg/fusionauth/auth"
)

const testAPIKey string = "6b87a398-39f2-4692-927b-13188a81a9a3"

func TestCreateClient(t *testing.T) {
	client := FusionAuthClient{URL: "http://localhost/9011", APIKey: testAPIKey}
	if client.APIKey != testAPIKey {
		t.Errorf("FusionAuthClient.APIKey = %s, want %s", client.APIKey, testAPIKey)
	}
}
