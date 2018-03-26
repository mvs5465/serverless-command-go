package main

import (
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func sendRequest(conf configuration) string {
	log.Info("Executing rest request: %s\n", conf.RemoteUrl)

	req, reqErr := http.NewRequest(conf.HttpMethod, conf.RemoteUrl, nil)
	req.Header.Set("Content-Type", "application/json")

	if reqErr != nil {
		log.Error(reqErr.Error())
		return "Error"
	}

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	res, resErr := spaceClient.Do(req)

	if resErr != nil {
		log.Info(resErr.Error())
		return "Error"
	}

	if conf.CheckResponseCode != "false" {
		log.Info("Checking response code...")
		if res.StatusCode == http.StatusOK {
			log.Info("Request successful. Http Status: %s", res.StatusCode)
		} else {
			log.Error("Request error. Http Status: %s", res.StatusCode)
		}
	} else {
		log.Info("Not checking response code.")
	}

	if res != nil {
		bodyBytes, err2 := ioutil.ReadAll(res.Body)
		bodyString := string(bodyBytes)
		if err2 != nil {
			log.Error("Error: %s", err2)
		} else {
			log.Info("Response code error: %s", res.StatusCode)
		}

		if bodyString != "" {
			log.Info("Message contains body")
		}
	}
	return "Success"
}
