package main

import (
	"github.com/urfave/cli"
)

// configAddFlags - Flags supported by the cal config add command.
var (
	configAddFlags = []cli.Flag{
		cli.StringFlag{
			Name:  "secret, s",
			Usage: "Provide path to clientId.json file.",
		},
	}
)

// configAddCmd - Command supporting adding a calendar to the current config.
var configAddCmd = cli.Command{
	Name:   "add",
	Usage:  "Add a new Google Calendar to the current config.",
	Flags:  append(configAddFlags, globalCalFlags...),
	Action: mainAddCal,
}

// mainAddCal - controls flow of the config add command.
// Responsible for argument parsing and then execution of the command.
func mainAddCal(ctx *cli.Context) error {
	if len(ctx.Args()) == 0 {
		// Requires user to enter alias of calendar being added.
		cli.ShowCommandHelp(ctx, ctx.Args().First())
		return nil
	}
	alias := ctx.Args()[0]
	secretPath := ctx.String("secret")
	return addCal(alias, secretPath)
}

// addCal - adds a new calendar to the config.
// Currently only supports Google Calendar.
func addCal(alias, secretPath string) error {
	// Load existing config.
	cfg, err := loadCalConfig()
	if err != nil {
		return err
	}
	// The to be stored calendar.
	newCal := calendar{
		Alias: alias,
	}

	// If path to .json is not provided assume oAuth is unnecessary.
	if secretPath != "" {
		// Retrieve a new oauth for the calendar.
		token, oauthConfig, err := getOAuth(alias, secretPath)
		if err != nil {
			return err
		}
		newCal.Token = token
		newCal.OAuthConfig = oauthConfig
	}
	// Store and save newly created calendar.
	cfg.addCalendar(newCal)
	saveCalConfig(cfg)
	return nil
}
