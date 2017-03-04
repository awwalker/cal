package main

import (
	"github.com/urfave/cli"
)

// configRemoveCmd - Command supporting removing a calendar to the current config.
var configRemoveCmd = cli.Command{
	Name:   "remove",
	Usage:  "Remove a Google Calendar from the current config.",
	Action: mainRemoveCal,
}

// mainRemoveCal - controls flow of the config remove command.
// Responsible for argument parsing and then execution of the command.
func mainRemoveCal(ctx *cli.Context) error {
	if len(ctx.Args()) == 0 {
		// Requires user to enter alias of calendar being removed.
		println("Remove command requires alias argument. ")
		cli.ShowCommandHelp(ctx, ctx.Args().First())
		return nil
	}
	alias := ctx.Args()[0]
	return removeCal(alias)
}

// removeCal - removes a calendar from the config.
// Currently only supports Google Calendar.
func removeCal(alias string) error {
	// Load existing config.
	cfg, err := loadCalConfig()
	if err != nil {
		return err
	}

	if _, ok := cfg.Calendars[alias]; ok {
		delete(cfg.Calendars, alias)
		println("Calendar successfully removed")
	} else {
		println("Calendar " + alias + " does not exist")
	}

	// Save newly updated config file.
	saveCalConfig(cfg)
	return nil
}
