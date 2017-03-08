package main

import "github.com/urfave/cli"

// configFlags
var (
	configFlags = []cli.Flag{}
)

// configCmd - Base config manipulation command.
var configCmd = cli.Command{
	Name:   "config",
	Usage:  "Manage cal configuration file.",
	Action: mainConfig,
	Flags:  append(configFlags, globalCalFlags...),
	Subcommands: []cli.Command{
		configAddCmd,
		// configRemoveCmd,
	},
}

// TODO: Fill in just show help here. If config is called
// without a subcommand.
func mainConfig(ctx *cli.Context) error {
	cli.ShowCommandHelp(ctx, ctx.Args().First())
	return nil
}
