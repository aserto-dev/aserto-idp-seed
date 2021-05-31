package cmdx

import (
	"github.com/aserto-dev/aserto-idp-seed/pkg/auth0"
	"github.com/aserto-dev/aserto-idp-seed/pkg/config"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

// SeedCommand -- seed, add entries from the seed file.
func SeedCommand() *cli.Command {
	return &cli.Command{
		Name:  "seed",
		Usage: "seed users",
		Flags: []cli.Flag{
			InputFileFlag(),
			SpewFlag(),
			NoCountFlag(),
			DryrunFlag(),
			UserMetadataFlag(),
			AppMetadataFlag(),
		},
		Before: configRetriever,
		Action: seedHandler,
	}
}

func seedHandler(c *cli.Context) (err error) {
	// get config from context
	cfg := config.FromContext(c.Context)

	if err := cfg.Auth0.Validate(); err != nil {
		return errors.Wrapf(err, "auth0 configuration missing")
	}

	if err := cfg.TemplateParams.Validate(); err != nil {
		return errors.Wrapf(err, "template values missing")
	}

	mgr := auth0.NewManager(cfg.Auth0)

	if err := mgr.Init(); err != nil {
		return err
	}

	mgr.Dryrun(c.Bool(flagDryRun))
	mgr.Spew(c.Bool(flagSpew))
	mgr.NoCount(c.Bool(flagNoCount))
	mgr.ImportUserMetadata((c.Bool(flagUserMetadata)))
	mgr.ImportAppMetadata((c.Bool(flagAppMetadata)))

	if err := mgr.Seed(cfg.TemplateParams); err != nil {
		return err
	}

	return nil
}
