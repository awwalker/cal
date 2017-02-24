package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var addCalCmd = cli.Command{
	Name:   "addCal",
	Usage:  "Configure new calendar",
	Action: addCal,
}

func addCal() (bool, error) {
	cfg, err := loadCalConfig()
	if err != nil {
		return false, err
	}

	reader := bufio.NewReader(os.Stdin)
	newCal := new(calendar)

	fmt.Println("Enter calendar alias: ")
	aliasTemp, _ := reader.ReadString('\n')
	newCal.Alias = aliasTemp

	// TODO: repeat for all other fields in calendar struct

	cfg.addCalendar(*newCal)
	saveCalConfig(cfg)

	return true, nil
}
