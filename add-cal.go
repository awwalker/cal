package main

import (
	"github.com/urfave/cli"
)

var configCmd = cli.Command{
	Name:   "config",
	Usage:  "Configure calendar setup",
	Action: mainConfig,
}

var (
	addCalFlags = []cli.Flag{
		cli.StringFlag{
			Name:  "secret, s",
			Usage: "Path to clientId.json",
		},
	}
)

func mainConfig(ctx *cli.Context) error {
	//TODO: Check arguments
	args := ctx.Args()

	command := args[0]
	alias := args[1]

	secret := ctx.String("secret")

	if command == "add" {
		addCal(alias, secret)
	} else if command == "rm" {
		//removeCal(alias)
	}
	return nil
}

func addCal(alias, secretPath string) error {
	cfg, err := loadCalConfig()
	if err != nil {
		return err
	}

	newCal := new(calendar)
	newCal.Alias = alias

	if secretPath != "" {
		token, oauthConfig, err := getOauth(alias, secretPath)
		if err == nil {
			newCal.Token = token
			newCal.OAuthConfig = oauthConfig
		}
	}

	cfg.addCalendar(*newCal)
	saveCalConfig(cfg)

	return nil
}
