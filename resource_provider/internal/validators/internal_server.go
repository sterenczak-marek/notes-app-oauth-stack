package validators

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/sterenczak-marek/notes-app-oauth-stack/resource_provider/config"
)

func validateInternalServerToken(token string) bool {
	client := http.DefaultClient

	validateData := config.OAuthValidateData["internal"]

	tokenData := url.Values{
		"token": []string{token},
	}
	req, err := http.NewRequest(http.MethodPost, validateData["URL"], strings.NewReader(tokenData.Encode()))
	if err != nil {
		log.Fatalf("unable to create validate request: %s", err)
	}
	req.SetBasicAuth(validateData["ClientID"], validateData["ClientSecret"])
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("error on validating token: %s", err)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		log.Panicf("Client authorization failed on OAuth server.")
	} else if resp.StatusCode != http.StatusOK {
		log.Panicf("Unexpected status code result. Accepts only 200, received: %d", resp.StatusCode)
	}

	var content map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Unable to read response body: %s", err)
	}
	resp.Body.Close()
	err = json.Unmarshal(body, &content)
	if err != nil {
		log.Fatalf("Unable to parse JSON response: %s", err)
	}

	return content["active"].(bool)
}

func init() {
	OAuthValidators.register("internal", validateInternalServerToken)
}
