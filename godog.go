package main

import (
	"errors"
	"os"
	"strings"

	"github.com/twistingmercury/godog/commands"
)

var buildDate = "{not set}"
var buildVersion = "{not set}"
var buildCommit = "{not set}"
var buildConfig = "debug"

func main() {
	if err := validateCfg(); err != nil {
		println(err.Error())
		os.Exit(2)
	}

	commands.SetBuildInfo(buildDate, buildVersion, buildCommit)
	if err := commands.Execute(); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func validateCfg() (err error) {
	const (
		api  = "missing value for DD_API_KEY"
		app  = "missing value for DD_APP_KEY"
		site = "missing value for DD_SITE"
	)

	sb := strings.Builder{}

	msg := func(s string) {
		if sb.Len() == 0 {
			sb.WriteString(s)
		} else {
			sb.WriteString("; " + s)
		}
	}

	if len(os.Getenv("DD_API_KEY")) == 0 {
		msg(api)
	}
	if len(os.Getenv("DD_APP_KEY")) == 0 {
		msg(app)
	}
	if len(os.Getenv("DD_SITE")) == 0 {
		msg(site)
	}

	if sb.Len() > 0 {
		err = errors.New(sb.String())
	}

	return
}
