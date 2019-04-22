package main

import (
	"github.com/sirupsen/logrus"
	//logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
	//"gopkg.in/gemnasium/logrus-airbrake-hook.v2"
	//"log/syslog"
	"github.com/multiplay/go-slack/chat"
	"github.com/multiplay/go-slack/lrhook"
)

func Bearychar() {
	cfg := lrhook.Config{
		MinLevel: logrus.InfoLevel,
		Message: chat.Message{
			Text: "Error Occur",
		},
		Attachment: chat.Attachment{
			Title: "hjiang.txt",
		},
	}

	h := lrhook.New(cfg, "https://hook.bearychat.com/=bw9Y1/incoming/********")
	logrus.SetFormmatter(&logrus.JSONFommater{})
	logrus.AddHook(h)
	logrus.Warn("")
}

func init() {
	log.AddHook(airbrake.NewHook(123, "xyz", "production"))

	hook, err := logrus_syslog.NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, "")
	if err != nil {
		log.Error("Unable to connect to local syslog deamon")
	} else {
		log.AddHook(hook)
	}
}
