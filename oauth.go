package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	gCalendar "google.golang.org/api/calendar/v3"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// tokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func tokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// saveToken - save a token to the config file under the calendar it belongs to.
func saveToken(token *oauth2.Token, alias string) error {
	cfg, err := loadCalConfig()
	if err != nil {
		return err
	}
	// Update the config with token.
	if _, ok := cfg.Calendars[alias]; !ok {
		return nil
	}
	curCalendar := cfg.Calendars[alias]
	curCalendar.Token = token
	saveCalConfig(cfg)
	return nil
}

// tokenFromConfig - returned a cached token if one exists.
func tokenFromConfig(alias string) (*oauth2.Token, error) {
	cfg, err := loadCalConfig()
	if err != nil {
		return nil, err
	}
	// If a token exists.
	if cfg.Calendars[alias].Token != nil {
		return cfg.Calendars[alias].Token, nil
	}
	return nil, fmt.Errorf("No Token in Config.")
}

// getToken - returns an oauth token from either a cache or instantiates a new one.
func getToken(alias string) (*oauth2.Token, error) {
	oAuthToken, err := tokenFromConfig(alias)
	if err != nil {
		// Assumes OAuth Config is stored in ~/.cal/client_id.json
		calConfigDir, err := getCalConfigDir()
		if err != nil {
			return nil, err
		}
		clientID := filepath.Join(calConfigDir, globalCalClientIDFile)
		clientBytes, err := ioutil.ReadFile(clientID)
		if err != nil {
			return nil, err
		}
		oAuthConfig, err := google.ConfigFromJSON(clientBytes, gCalendar.CalendarScope)
		if err != nil {
			return nil, err
		}
		oAuthToken = tokenFromWeb(oAuthConfig)
		// Only need to cache here.
		if err := saveToken(oAuthToken, alias); err != nil {
			return nil, err
		}
	}
	return oAuthToken, nil
}
