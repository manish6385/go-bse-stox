package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Initialise logrus configuration
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(log.TraceLevel)
	log.SetOutput(os.Stdout)
}
func main() {

	log.Println("Hello World!")

}
