package main

import (
	"log"

	"github.com/urfave/cli"
)

var calHelpTemplate = ``
var calFlags = []cli.Flag{}

func main() {
	if err := initCal(); err != nil {
		log.Fatalf("oops: %v\n", err)
	}
	app := registerApp()
	app.RunAndExitOnError()
}

func registerApp() *cli.App {
	registerCmd(configCmd)

	app := cli.NewApp()
	app.Usage = "Manage calendar from the Command Line"
	app.Flags = append(calFlags, globalCalFlags...)
	app.Version = "0.0.1"
	app.Commands = commands
	app.CommandNotFound = func(ctx *cli.Context, command string) {
		cli.ShowAppHelp(ctx)
	}
	return app
}

// initCal - run on startup to initialize config.
func initCal() error {
	if !isCalConfigExists() {
		if err := saveCalConfig(newConfig()); err != nil {
			return err
		}
	}
	return nil
}

var commands = []cli.Command{}

func registerCmd(cmd cli.Command) {
	commands = append(commands, cmd)
}
