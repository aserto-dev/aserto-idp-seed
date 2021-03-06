package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aserto-dev/aserto-idp-seed/cmd/aserto-idp-seed/cmdx"
	"github.com/aserto-dev/aserto-idp-seed/pkg/version"

	_ "github.com/joho/godotenv/autoload"

	"github.com/urfave/cli/v2"
)

const (
	appName  = "aserto-idp-seed"
	appUsage = ""
)

func main() {
	l, err := os.Create("./aserto-idp-seed.log")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	log.SetOutput(l)
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("aserto-idp-seed")

	appl := cli.NewApp()
	appl.EnableBashCompletion = true
	appl.Name = appName
	appl.Usage = appUsage
	appl.HideHelp = true
	appl.HideHelpCommand = true
	appl.HideVersion = true
	appl.Version = version.GetInfo().String()
	appl.Commands = []*cli.Command{
		cmdx.SeedCommand(),
		cmdx.ResetCommand(),
		cmdx.UsersCommand(),
		cmdx.ValidateCommand(),
		cmdx.VersionCommand(),
	}

	ctx := context.Background()

	if err := appl.RunContext(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error %v\n", err)
	}
}
