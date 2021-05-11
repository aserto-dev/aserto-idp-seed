package cmdx

import (
	"github.com/aserto-dev/aserto-idp-seed/pkg/auth0"
	"github.com/aserto-dev/aserto-idp-seed/pkg/config"
	"github.com/urfave/cli/v2"
)

// SeedCommand -- seed, add entries from the seed file.
func SeedCommand() *cli.Command {
	return &cli.Command{
		Name:  "seed",
		Usage: "seed",
		Flags: []cli.Flag{
			InputFileFlag(),
			SpewFlag(),
			DryrunFlag(),
		},
		Before: configRetriever,
		Action: seedHandler,
	}
}

func seedHandler(c *cli.Context) (err error) {
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

	if err := mgr.Seed(cfg.TemplateParams); err != nil {
		return err
	}

	return nil
}
