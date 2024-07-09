package infisical

import (
	"os"
	"testing"
	"time"

	infisical "github.com/infisical/go-sdk"
)

// For this test to work, set access token ttl to 10 seconds in the Infisical dashboard.
func TestTokenRefresh(t *testing.T) {
	clientID := os.Getenv("INFISICAL_CLIENT_ID")
	clientSecret := os.Getenv("INFISICAL_CLIENT_SECRET")
	projectID := os.Getenv("INFISICAL_PROJECT_ID")

	if clientID == "" || clientSecret == "" || projectID == "" {
		t.Skip("Skipping test because INFISICAL_CLIENT_ID, INFISICAL_CLIENT_SECRET or INFISICAL_PROJECT_ID are not set")
	}

	client := NewClient(AccessTokenRenewer(UniversalAuthLogin(clientID, clientSecret), WithExpirationWindow(10*time.Second)), WithSiteURL(os.Getenv("INFISICAL_SITE_URL")))

	_, err := client.Secrets().Delete(infisical.DeleteSecretOptions{
		ProjectID:   projectID,
		Environment: "dev",
		SecretKey:   "infisical-go-test",
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Secrets().Create(infisical.CreateSecretOptions{
		ProjectID:   projectID,
		Environment: "dev",
		SecretKey:   "infisical-go-test",
		SecretValue: "foo",
	})
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(30 * time.Second)

	_, err = client.Secrets().Retrieve(infisical.RetrieveSecretOptions{
		ProjectID:   projectID,
		Environment: "dev",
		SecretKey:   "infisical-go-test",
	})
	if err != nil {
		t.Fatal(err)
	}
}
