package main

import (
	"log"
	"log/syslog"
)

func main() {
	// Create a new logger
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "pensu")
	if err != nil {
		log.Fatal()
	} else {
		log.SetOutput(sysLog)
	}

	log.Panic(sysLog)
}
