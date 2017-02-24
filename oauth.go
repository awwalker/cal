package main

import (
	"fmt"
	"io/ioutil"
	"log"

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

// saveOAuth - save a token to the config file under the calendar it belongs to.
func saveOAuth(token *oauth2.Token, config *oauth2.Config, alias string) error {
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
	curCalendar.OAuthConfig = config
	saveCalConfig(cfg)
	return nil
}

// tokenFromConfig - returned a cached token if one exists.
func oAuthFromConfig(alias string) (*oauth2.Token, *oauth2.Config, error) {
	cfg, err := loadCalConfig()
	if err != nil {
		return nil, nil, err
	}
	// If alias exists.
	cal, ok := cfg.Calendars[alias]
	if !ok {
		return nil, nil, fmt.Errorf("No Alias in Config Matching: %s", alias)
	}
	// If a token exists.
	if cal.Token != nil && cal.OAuthConfig != nil {
		return cal.Token, cal.OAuthConfig, nil
	}
	return nil, nil, fmt.Errorf("Could not get OAuth information")
}

// getOAuth - sets an oauth token from either a cache or instantiates a new one for a calendar.
func getOAuth(alias, clientPath string) (*oauth2.Token, *oauth2.Config, error) {
	oAuthToken, oAuthConfig, err := oAuthFromConfig(alias)
	if err != nil {
		// Assumes OAuth Config is stored in clientPath.
		clientBytes, err := ioutil.ReadFile(clientPath)
		if err != nil {
			return nil, nil, err
		}
		oAuthConfig, err := google.ConfigFromJSON(clientBytes, gCalendar.CalendarScope)
		if err != nil {
			return nil, nil, err
		}
		oAuthToken = tokenFromWeb(oAuthConfig)
	}
	fmt.Printf("%#v, %#v", oAuthToken, oAuthConfig)
	return oAuthToken, oAuthConfig, nil
}
