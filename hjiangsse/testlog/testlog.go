package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	//log as JSON instead of the deafault ASCII formatter
	log.SetFormatter(&log.JSONFormatter{})

	//Output to stdout instead of the default stderr
	//Can be any io.wirter
	log.SetOutput(os.Stdout)

	//Only log the warning serverity or above
	//log.SetLevel(log.WarnLevel)
	//log.SetLevel(log.InfoLevel)
}

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Warn("The ice breaks!")

	//fmt.Println("This is a test line.")

	//reuse fields between loging statements by re-using
	//the logrus.Entry returned from WithFields
	contextLogger := log.WithFields(log.Fields{
		"common": "This is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
