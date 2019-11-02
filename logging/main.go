package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	// Get the name of current program
	programName := filepath.Base(os.Args[0])
	fmt.Println(programName)

	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if err != nil {
		log.Fatal()
	} else {
		log.SetOutput(sysLog)
	}

	log.Println("LOG_INFO+LOG_LOCAL7: Logging in Go")
}
