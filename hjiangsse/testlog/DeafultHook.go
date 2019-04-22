package main

import (
	"github.com/sirupsen/logrus"
)

//A user defined hook
type Hook struct {
	name string
}

func NewHook(name string) (h *Hook) {
	return &Hook{name}
}

func (hook *Hook) Fire(entry *logrus.Entry) (err error) {
	entry.Data["appName"] = "MyApp"
	return nil
}

func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func init() {
	logrus.AddHook(NewHook("test"))
}

func main() {
	logrus.WithFields(logrus.Fields{
		"animal": "dog",
		"age":    "10",
	}).Info("This dog is ten years old now!")
}
