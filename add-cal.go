package main 

import (
	"fmt"
	"os"
	"bufio"
)

var configCalCmd = cli.Command {
	Name: "addCal",
	Usage: "Configure new calendar",
}

func addCal() (bool, error){
	cfg, err := loadConfig()
	if err != nil {
		return false, err
	}

	reader := bufio.NewReader(os.Stdin)
	newCal := new(calendar)

	fmt.Println("Enter calendar alias: ")
	newCal.Alias, _ := reader.ReadString('\n')

	// TODO: repeat for all other fields in calendar struct

	cfg.addCalendar(&newCal)
	saveCalConfig(cfg)
}
