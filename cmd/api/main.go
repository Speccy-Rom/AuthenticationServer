package main

import (
	log "github.com/sirupsen/logrus"
	//"log"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}


