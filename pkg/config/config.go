package config

import (
	"context"
	"os"

	"github.com/pkg/errors"
)

const (
	configKey            = "config"
	EnvAuth0Domain       = "AUTH0_DOMAIN"
	EnvAuth0ClientID     = "AUTH0_CLIENT_ID"
	EnvAuth0ClientSecret = "AUTH0_CLIENT_SECRET" //nolint: gosec
	EnvTemplCorporation  = "TEMPL_CORPORATION"
	EnvTemplEmailDomain  = "TEMPL_EMAIL_DOMAIN"
	EnvTemplPassword     = "TEMPL_PASSWORD"
)

type key string

// Config - config structure.
type Config struct {
	Auth0          *Auth0
	TemplateParams *TemplateParams
}

// Auth0 - Auth0 config structure.
type Auth0 struct {
	Domain       string
	ClientID     string
	ClientSecret string
}

const errEnvNotSetMsg = "%s environment variable not set"

func (a *Auth0) Validate() error {
	switch {
	case a.Domain == "":
		return errors.Errorf(errEnvNotSetMsg, EnvAuth0Domain)
	case a.ClientID == "":
		return errors.Errorf(errEnvNotSetMsg, EnvAuth0ClientID)
	case a.ClientSecret == "":
		return errors.Errorf(errEnvNotSetMsg, EnvAuth0ClientSecret)
	}
	return nil
}

type TemplateParams struct {
	Corporation string
	EmailDomain string
	Password    string
}

func (t *TemplateParams) Validate() error {
	switch {
	case t.Corporation == "":
		return errors.Errorf(errEnvNotSetMsg, EnvTemplCorporation)
	case t.EmailDomain == "":
		return errors.Errorf(errEnvNotSetMsg, EnvTemplEmailDomain)
	case t.Password == "":
		return errors.Errorf(errEnvNotSetMsg, EnvTemplPassword)
	}
	return nil
}

// FromEnv - create config instance from environment variables
func FromEnv() *Config {
	cfg := Config{
		Auth0: &Auth0{
			Domain:       os.Getenv(EnvAuth0Domain),
			ClientID:     os.Getenv(EnvAuth0ClientID),
			ClientSecret: os.Getenv(EnvAuth0ClientSecret),
		},
		TemplateParams: &TemplateParams{
			Corporation: os.Getenv(EnvTemplCorporation),
			EmailDomain: os.Getenv(EnvTemplEmailDomain),
			Password:    os.Getenv(EnvTemplPassword),
		},
	}
	return &cfg
}

// Key -- context key for config.
func Key() interface{} {
	var k = key(configKey)
	return k
}

// FromContext -- extract config from context value.
func FromContext(ctx context.Context) *Config {
	cfg, ok := ctx.Value(Key()).(*Config)
	if !ok {
		return nil
	}

	return cfg
}
