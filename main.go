package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/scottdware/go-junos"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := mainInternal()
	if err != nil {
		log.Errorf("ERROR: %v", err)
		os.Exit(1)
	}
	log.Info("DONE.")
}

func mainInternal() error {
	log.SetLevel(log.DebugLevel)

	// gather data
	juniperHost := os.Getenv("JUNIPER_HOST")
	if juniperHost == "" {
		return fmt.Errorf("Missing environment variable: JUNIPER_HOST")
	}

	juniperUser := os.Getenv("JUNIPER_USER")
	if juniperUser == "" {
		return fmt.Errorf("Missing environment variable: JUNIPER_USER")
	}

	juniperPassword := os.Getenv("JUNIPER_PASSWORD")
	if juniperPassword == "" {
		return fmt.Errorf("Missing environment variable: JUNIPER_PASSWORD")
	}

	juniperCommand := os.Getenv("JUNIPER_COMMAND")
	if juniperCommand == "" {
		return fmt.Errorf("Missing environment variable: JUNIPER_COMMAND")
	}

	// perform login
	log.Infof("Connecting to %v...", juniperHost)

	auth := &junos.AuthMethod{
		Credentials: []string{juniperUser, juniperPassword},
	}

	jnpr, err := junos.NewSession(juniperHost, auth)
	if err != nil {
		return err
	}

	defer jnpr.Close()

	// request disconnect
	commands := strings.Split(juniperCommand, ";;")
	for _, cmd := range commands {
		log.Infof("Running command: %v", cmd)
		res, err := jnpr.Command(cmd, "text")
		if err != nil {
			if err.Error() == "EOF" {
				log.Infof("  Result: <none>")
				continue
			}
			return err
		}

		log.Infof("  Result:\n%v", strings.TrimSpace(res))
	}

	return nil
}
