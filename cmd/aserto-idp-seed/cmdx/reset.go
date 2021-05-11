package cmdx

import (
	"github.com/aserto-dev/aserto-idp-seed/pkg/auth0"
	"github.com/aserto-dev/aserto-idp-seed/pkg/config"
	"github.com/urfave/cli/v2"
)

// ResetCommand -- reset state, removes all entries added from the seed file.
func ResetCommand() *cli.Command {
	return &cli.Command{
		Name:  "reset",
		Usage: "reset",
		Flags: []cli.Flag{
			InputFileFlag(),
			SpewFlag(),
			DryrunFlag(),
		},
		Before: configRetriever,
		Action: resetHandler,
	}
}

func resetHandler(c *cli.Context) (err error) {
	// get config from context
	cfg := config.FromContext(c.Context)

	mgr := auth0.NewManager(
		cfg.Auth0,
		c.Path(flagInputFile),
	)

	if err := mgr.Init(); err != nil {
		return err
	}

	mgr.Dryrun(c.Bool(flagDryRun))
	mgr.Spew(c.Bool(flagSpew))

	if err := mgr.Reset(); err != nil {
		return err
	}

	return nil
}
