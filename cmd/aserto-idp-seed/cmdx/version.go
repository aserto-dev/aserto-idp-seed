package cmdx

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const (
	versionName  = "version"
	versionUsage = "version information"
)

// VersionCommand -- version command definition.
func VersionCommand() *cli.Command {
	return &cli.Command{
		Name:     versionName,
		Usage:    versionUsage,
		HideHelp: true,
		Action: func(c *cli.Context) (err error) {
			fmt.Fprintf(c.App.Writer, "%s - %s\n",
				c.App.Name,
				c.App.Version,
			)
			return nil
		},
	}
}
