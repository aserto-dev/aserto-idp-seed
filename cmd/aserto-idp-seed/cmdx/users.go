package cmdx

import (
	"github.com/aserto-dev/aserto-idp-seed/pkg/auth0"
	"github.com/aserto-dev/aserto-idp-seed/pkg/config"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

// UsersCommand -- list existing users.
func UsersCommand() *cli.Command {
	return &cli.Command{
		Name:  "users",
		Usage: "list existing users",
		Flags: []cli.Flag{
			SpewFlag(),
			NoCountFlag(),
		},
		Before: configRetriever,
		Action: usersHandler,
	}
}

func usersHandler(c *cli.Context) (err error) {
	// get config from context
	cfg := config.FromContext(c.Context)

	if err := cfg.Auth0.Validate(); err != nil {
		return errors.Wrapf(err, "auth0 configuration missing")
	}

	mgr := auth0.NewManager(cfg.Auth0)

	if err := mgr.Init(); err != nil {
		return err
	}

	mgr.Spew(c.Bool(flagSpew))
	mgr.NoCount(c.Bool(flagNoCount))

	if err := mgr.Users(); err != nil {
		return err
	}

	return nil
}
