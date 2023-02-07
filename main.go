package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/decentralgabe/go-rdr-client/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logrus.WithError(err).Error("failed to execute command")
		os.Exit(1)
	}
}
