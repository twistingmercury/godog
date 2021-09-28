package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/twistingmercury/godog/commands"
)

var buildDate = "{not set}"
var buildVersion = "{not set}"
var buildCommit = "{not set}"
var buildConfig = "debug"

func main() {
	if strings.ToLower(buildConfig) == "debug" {
		if err := godotenv.Load("./godog.env"); err != nil {
			log.Fatal("error loading .env file")
		}
	}

	version()

	if err := commands.Execute(); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func version() {
	if os.Args[1] == "--version" {
		commands.Logo()
		fmt.Printf("\n- Version %s\n- Build Date: %s\n- Commit: %s\n- Build Configuration: %s\n", buildVersion, buildDate, buildCommit, buildConfig)
		os.Exit(0)
	}
}
