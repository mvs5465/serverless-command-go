package main

import (
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

//// Checklist
// ( x ) JSON LOGS
// ( x ) LOG level
// ( x ) dry run
// ( x ) -- if dry run --> env var? --> dont do anything, just log what you would have done
// ( x ) Find out if it is a URL or bash command?
// (  ) add a field for headers
// (  ) validate ssl certificaties (skip for now)

//////

func setLogLevel(level string) {
	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warning":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

func (r *yamlConfig) parseFile(yamlFile []byte) *yamlConfig {
	err := yaml.Unmarshal(yamlFile, &r)
	if err != nil {
		log.Info(err)
	}
	return r
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	setLogLevel(os.Getenv("LOG_LEVEL"))

	var dryRun string = os.Getenv("DRY_RUN")
	if dryRun != "true" {
		log.Info("Not DRY_RUNning")
	} else {
		log.Info("DRY_RUNning")
	}

	yamlFile, err := ioutil.ReadFile("./files/sample2.yaml")

	var r yamlConfig
	if err != nil {
		log.Info(err)
	} else {
		r.parseFile(yamlFile)
	}

	for _, cmd0 := range r.Descriptions[1].Commands {
		if cmd0.Configuration.Command != "" {
			// assume its a command
			log.Info("Executing command...")
			var resultCmd string = executeCommand(cmd0.Configuration)
			log.Info("Result: " + resultCmd)
		} else {
			// assume its a rest request
			log.Info("Executing rest request")
			var resultRr string = sendRequest(cmd0.Configuration)
			log.Info("Result:  " + resultRr)
		}
	}
}
