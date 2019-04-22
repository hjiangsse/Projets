package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

//create a new instance of the logger
var log = logrus.New()

func main() {
	//log.Out = os.Stdout

	//write log into a file
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}
