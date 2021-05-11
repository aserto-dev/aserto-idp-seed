package cmdx

import (
	"context"

	"github.com/aserto-dev/aserto-idp-seed/pkg/config"

	"github.com/urfave/cli/v2"
)

// configRetriever -- pre-command handler
// loads config and persists config in context
func configRetriever(c *cli.Context) error {
	cfg := config.FromEnv()

	c.Context = context.WithValue(c.Context, config.Key(), cfg)

	return nil
}
