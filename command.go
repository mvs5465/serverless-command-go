package main

import (
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

func executeCommand(conf configuration) string {
	log.Info("Executing command: %s\n", conf.Command)
	args := strings.Fields(conf.Command)
	cmd := exec.Command(args[0], args[1:]...)
	_, cerr := cmd.Output()

	if cerr != nil {
		log.Error(cerr.Error())
		return "Error"
	} else {
		log.Info("Http request successful")
	}

	return "Success"
}
