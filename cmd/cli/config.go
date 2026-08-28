package main

import (
	"flag"
	"fmt"
	"io"
)

const (
	defaultServerAddr = "localhost:50051"
	defaultClientCert = "certs/client.crt"
	defaultClientKey  = "certs/client.key"
	defaultCACert     = "certs/ca.crt"
)

type cliConfig struct {
	ServerAddr     string
	ClientCertPath string
	ClientKeyPath  string
	CACertPath     string
}

func parseCLIConfig(args []string, lookupEnv func(string) (string, bool)) (cliConfig, []string, error) {
	config := cliConfig{
		ServerAddr:     envOrDefault(lookupEnv, "SENTRY_SERVER", defaultServerAddr),
		ClientCertPath: envOrDefault(lookupEnv, "SENTRY_CLIENT_CERT", defaultClientCert),
		ClientKeyPath:  envOrDefault(lookupEnv, "SENTRY_CLIENT_KEY", defaultClientKey),
		CACertPath:     envOrDefault(lookupEnv, "SENTRY_CA_CERT", defaultCACert),
	}

	flags := flag.NewFlagSet("cli", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	flags.StringVar(&config.ServerAddr, "server", config.ServerAddr, "Sentry server address")
	flags.StringVar(&config.ClientCertPath, "cert", config.ClientCertPath, "Client certificate path")
	flags.StringVar(&config.ClientKeyPath, "key", config.ClientKeyPath, "Client private key path")
	flags.StringVar(&config.CACertPath, "ca", config.CACertPath, "CA certificate path")
	if err := flags.Parse(args); err != nil {
		return cliConfig{}, nil, err
	}

	for name, path := range map[string]string{
		"client certificate": config.ClientCertPath,
		"client key":         config.ClientKeyPath,
		"CA certificate":     config.CACertPath,
	} {
		if path == "" {
			return cliConfig{}, nil, fmt.Errorf("%s path cannot be empty", name)
		}
	}

	return config, flags.Args(), nil
}

func envOrDefault(lookupEnv func(string) (string, bool), name, fallback string) string {
	if value, ok := lookupEnv(name); ok && value != "" {
		return value
	}
	return fallback
}
