// Copyright (c) 2018-present, Giuseppe Lo Brutto All rights reserved
// Package main provides the command to make remote connection
// by a configuration hidden file name .remote_connections
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/gambledor/golang/remote-connector/file"
)

const (
	CONF_FILE_NAME string = ".remote_connections"
)

// remoteConnector execute the ssh connection to the chosen remoteMachine
func remoteConnector(remoteMachine file.RemoteMachine) {
	var connectionString string = fmt.Sprintf("%s@%s", remoteMachine.User, remoteMachine.Host)
	cmd := exec.Command("ssh", connectionString)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func main() {
	// 1. read configuration file for remote connections
	var remoteMachines []file.RemoteMachine

	remoteMachines, err := file.ReadConfigFile(os.Getenv("HOME"), CONF_FILE_NAME)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// 2. show a menu of remote connection
	for idx, machine := range remoteMachines {
		fmt.Printf("%d) %s\n", idx, machine.Name)
	}
	fmt.Print("> ")

	// 3. the user makes a choise to witch machine wants to connect to
	var choise int
	if _, err := fmt.Scanf("%d", &choise); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("You've chosen to connect to", remoteMachines[choise].Host)

	// 4. make a ssh connction to the chosen machine
	var remoteMachine file.RemoteMachine = remoteMachines[choise]
	remoteConnector(remoteMachine)
}
