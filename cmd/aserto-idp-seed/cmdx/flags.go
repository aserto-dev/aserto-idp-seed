package cmdx

import (
	"github.com/aserto-dev/aserto-idp-seed/pkg/config"
	"github.com/urfave/cli/v2"
)

// InputFileFlag -- input file flag.
func InputFileFlag() cli.Flag {
	return &cli.PathFlag{
		Name:     flagInputFile,
		Aliases:  []string{},
		Usage:    usageInputFile,
		Required: false,
		Hidden:   true,
	}
}

// DryrunFlag -- dryrun flag.
func DryrunFlag() cli.Flag {
	return &cli.BoolFlag{
		Name:    flagDryRun,
		Aliases: []string{},
		Usage:   usageDryRun,
	}
}

// SpewFlag -- spew flag.
func SpewFlag() cli.Flag {
	return &cli.BoolFlag{
		Name:    flagSpew,
		Aliases: []string{},
		Usage:   usageSpew,
	}
}

// NoCountFlag -- no count flag.
func NoCountFlag() cli.Flag {
	return &cli.BoolFlag{
		Name:    flagNoCount,
		Aliases: []string{},
		Usage:   usageNoCount,
	}
}

func CorporationFlag() cli.Flag {
	return &cli.StringFlag{
		Name:    flagCorporation,
		Aliases: []string{},
		Usage:   usageCorporation,
		EnvVars: []string{config.EnvTemplCorporation},
	}
}

func EmailDomainFlag() cli.Flag {
	return &cli.StringFlag{
		Name:    flagEmailDomain,
		Aliases: []string{},
		Usage:   usageEmailDomain,
		EnvVars: []string{config.EnvTemplEmailDomain},
	}
}

func PasswordFlag() cli.Flag {
	return &cli.StringFlag{
		Name:    flagPassword,
		Aliases: []string{},
		Usage:   usagePassword,
		EnvVars: []string{config.EnvTemplPassword},
	}
}

func UserMetadataFlag() cli.Flag {
	return &cli.BoolFlag{
		Name:  flagUserMetadata,
		Usage: usageUserMetadata,
	}
}

func AppMetadataFlag() cli.Flag {
	return &cli.BoolFlag{
		Name:  flagAppMetadata,
		Usage: usageAppMetadata,
	}
}
