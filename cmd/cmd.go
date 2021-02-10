package cmd

import (
	// "zy/utils"

	"go-cli-starter/log"

	"github.com/urfave/cli/v2"
)

func runAction(c *cli.Context) error {

	message := c.String("message")
	log.Logger.Info(message)
	return nil
}

// Run Command
var Run = &cli.Command{
	Name:   "run",
	Usage:  "This is a command",
	Action: runAction,
	Flags: []cli.Flag{

		&cli.StringFlag{
			Name:    "message",
			Aliases: []string{"m"},
			Usage:   "Print Message.",
		},
	},
}
