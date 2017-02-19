package main

import "github.com/urfave/cli"

const (
	globalCalConfigDir    = ".cal/"
	globalCalConfigFile   = "config.json"
	globalCalClientIDFile = "client_id.json"
)

var (
	globalCalFlags = []cli.Flag{}
)
