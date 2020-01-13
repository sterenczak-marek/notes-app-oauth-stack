package validators

import (
	"log"
	"net/http"

	"github.com/sterenczak-marek/notes-app-oauth-stack/resource_provider/config"
)

func validateGithubToken(token string) bool {
	client := http.DefaultClient

	validateData := config.OAuthValidateData["github"]

	req, err := http.NewRequest(
		http.MethodGet,
		validateData["URL"],
		nil,
	)
	if err != nil {
		log.Fatalf("unable to create validate request: %s", err)
	}
	req.Header.Set(
		"Authorization",
		"Bearer "+token,
	)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("error on validating token: %s", err)
	}

	if resp.StatusCode == http.StatusOK {
		return true
	} else if resp.StatusCode == http.StatusUnauthorized {
		return false
	}

	log.Printf("Unexpected status code result. Accepts only 200, received: %d. Response: %v", resp.StatusCode, resp)
	return false
}

func init() {
	OAuthValidators.register("github", validateGithubToken)
}
