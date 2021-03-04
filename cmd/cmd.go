package cmd

import (
	// "zy/utils"

	// "go-cli-starter/log"

	"go-cli-starter/utils"

	"github.com/urfave/cli/v2"
)

func runAction(c *cli.Context) error {

	message := c.String("url")
	_, err := utils.HTTPGet(message)

	if err != nil {
		return nil
	}

	return nil
}

// Run Command
var Run = &cli.Command{
	Name:   "run",
	Usage:  "This is a command",
	Action: runAction,
	Flags: []cli.Flag{

		&cli.StringFlag{
			Name:    "url",
			Aliases: []string{"u"},
			Usage:   "Http Get url",
		},
	},
}
