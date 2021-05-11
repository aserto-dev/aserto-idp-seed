package cmdx

import (
	"github.com/aserto-dev/aserto-idp-seed/pkg/auth0"
	"github.com/aserto-dev/aserto-idp-seed/pkg/config"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

// ResetCommand -- reset state, removes all entries added from the seed file.
func ResetCommand() *cli.Command {
	return &cli.Command{
		Name:  "reset",
		Usage: "remove seeded users",
		Flags: []cli.Flag{
			InputFileFlag(),
			SpewFlag(),
			NoCountFlag(),
			DryrunFlag(),
		},
		Before: configRetriever,
		Action: resetHandler,
	}
}

func resetHandler(c *cli.Context) (err error) {
	// get config from context
	cfg := config.FromContext(c.Context)

	if err := cfg.Auth0.Validate(); err != nil {
		return errors.Wrapf(err, "auth0 configuration missing")
	}

	mgr := auth0.NewManager(cfg.Auth0)

	if err := mgr.Init(); err != nil {
		return err
	}

	mgr.Dryrun(c.Bool(flagDryRun))
	mgr.Spew(c.Bool(flagSpew))
	mgr.NoCount(c.Bool(flagNoCount))

	if err := mgr.Reset(); err != nil {
		return err
	}

	return nil
}
