package config

import (
	"context"
	"os"
)

const (
	configKey            = "config"
	envAuth0Domain       = "AUTH0_DOMAIN"
	envAuth0ClientID     = "AUTH0_CLIENT_ID"
	envAuth0ClientSecret = "AUTH0_CLIENT_SECRET" //nolint: gosec
	envTemplCorporation  = "TEMPL_CORPORATION"
	envTemplEmailDomain  = "TEMPL_EMAIL_DOMAIN"
	envTemplPassword     = "TEMPL_PASSWORD"
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

type TemplateParams struct {
	Corporation string
	EmailDomain string
	Password    string
}

// FromEnv - create config instance from environment variables
func FromEnv() *Config {
	cfg := Config{
		Auth0: &Auth0{
			Domain:       os.Getenv(envAuth0Domain),
			ClientID:     os.Getenv(envAuth0ClientID),
			ClientSecret: os.Getenv(envAuth0ClientSecret),
		},
		TemplateParams: &TemplateParams{
			Corporation: os.Getenv(envTemplCorporation),
			EmailDomain: os.Getenv(envTemplEmailDomain),
			Password:    os.Getenv(envTemplPassword),
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
