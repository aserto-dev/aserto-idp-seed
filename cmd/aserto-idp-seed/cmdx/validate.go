package cmdx

import (
	"fmt"

	"github.com/aserto-dev/aserto-idp-seed/pkg/auth0"
	"github.com/aserto-dev/aserto-idp-seed/pkg/config"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

// ValidateCommand -- validate ido connection.
func ValidateCommand() *cli.Command {
	return &cli.Command{
		Name:   "validate",
		Usage:  "validate idp connection",
		Flags:  []cli.Flag{},
		Before: configRetriever,
		Action: validateHandler,
	}
}

func validateHandler(c *cli.Context) (err error) {
	// get config from context
	cfg := config.FromContext(c.Context)

	fmt.Fprintf(c.App.Writer, ">>> validation configuration\n")
	if err := cfg.Auth0.Validate(); err != nil {
		return errors.Wrapf(err, "auth0 configuration missing")
	}

	mgr := auth0.NewManager(cfg.Auth0)
	if err := mgr.Validate(); err != nil {
		fmt.Fprintf(c.App.Writer, "!!! validation failed\n")
		return err
	}

	fmt.Fprintf(c.App.Writer, "+++ validation succeeded\n")
	return nil
}
