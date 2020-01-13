package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	DB                *gorm.DB
	OAuthValidateData = make(map[string]map[string]string)
)

func init() {
	initDB()
	getOAuthData("internal", "OAUTH_INTERNAL_SERVER_")
	getOAuthData("github", "OAUTH_GITHUB_")
}

func initDB() {
	dbPath := os.Getenv("DB")
	if dbPath == "" {
		log.Fatalf("provide `OAUTH_DB` env variable")
	}
	db, err := gorm.Open("postgres", dbPath)
	if err != nil {
		log.Fatalf("unable to connect to database: %s", err)
	}
	DB = db
}

func getOAuthData(providerName, prefix string) {
	var errors []string

	OAuthValidateData[providerName] = make(map[string]string)

	env := prefix + "VALIDATE_TOKEN_URL"
	errorMsg := "provide `%s` env variable"
	validateURL := os.Getenv(env)
	if validateURL == "" {
		errors = append(errors, fmt.Sprintf(errorMsg, env))
	} else {
		OAuthValidateData[providerName]["URL"] = validateURL
	}

	env = prefix + "VALIDATE_CLIENT_ID"
	clientID := os.Getenv(env)
	if clientID == "" {
		errors = append(errors, fmt.Sprintf(errorMsg, env))
	} else {
		OAuthValidateData[providerName]["ClientID"] = clientID
	}

	env = prefix + "VALIDATE_CLIENT_SECRET"
	clientSecret := os.Getenv(env)
	if clientID == "" {
		errors = append(errors, fmt.Sprintf(errorMsg, env))
	} else {
		OAuthValidateData[providerName]["ClientSecret"] = clientSecret
	}

	if len(errors) > 0 {
		log.Fatalf(strings.Join(errors, "\n"))
	}
}
