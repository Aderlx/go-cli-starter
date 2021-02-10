package main

import (
	"go-cli-starter/cmd"
	"go-cli-starter/config"
	"os"

	"github.com/urfave/cli/v2"
)

// cli app name
var appName = config.Get("app.Name").(string)

// cli app author name
var appAuthorName = config.Get("app.author.name").(string)

// cli app author email
var appAuthorEmail = config.Get("app.author.email").(string)

// cli app version
var appVersion = config.Get("app.version").(string)

// cli app usage
var appUsage = config.Get("app.usage").(string)

func main() {
	// fmt.Println(appAuthors)

	app := cli.NewApp()
	app.Name = appName
	app.Authors = []*cli.Author{
		{
			Name:  appAuthorName,
			Email: appAuthorEmail,
		},
	}
	app.Version = appVersion
	app.Usage = appUsage
	app.Commands = []*cli.Command{cmd.Run}
	// app.Flags
	app.Flags = append(app.Flags, cmd.Run.Flags...)

	app.Run(os.Args)
}
