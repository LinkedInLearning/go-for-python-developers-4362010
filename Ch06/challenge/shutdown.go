package main

import (
	"fmt"
	"log"
	"os"
)

var (
	logger = log.New(log.Writer(), "[shutdown ]", log.LstdFlags|log.Lshortfile)
)

func shutdownServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		logger.Printf("ERROR: can't open %q", pidFile)
		return err
	}
	defer file.Close()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		logger.Printf("ERROR: %q: bad pid - %s", pidFile, err)
		return err

	}

	logger.Printf("INFO: shutting down %d", pid)
	// TODO: Actual shutdown

	if err := os.Remove(pidFile); err != nil {
		log.Printf("WARNING: can't delete %q - %s", pidFile, err)
	}

	return nil
}

func main() {
	shutdownServer("app.pid")
}
