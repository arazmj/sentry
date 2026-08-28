package main

import (
	"flag"
	"fmt"
	"io"
	"strconv"
)

const (
	defaultPort       = 50051
	defaultServerCert = "certs/server.crt"
	defaultServerKey  = "certs/server.key"
	defaultCACert     = "certs/ca.crt"
)

type serverConfig struct {
	Port           int
	ServerCertPath string
	ServerKeyPath  string
	CACertPath     string
}

func parseServerConfig(args []string, lookupEnv func(string) (string, bool)) (serverConfig, error) {
	port, err := envPort(lookupEnv)
	if err != nil {
		return serverConfig{}, err
	}

	config := serverConfig{
		Port:           port,
		ServerCertPath: envOrDefault(lookupEnv, "SENTRY_SERVER_CERT", defaultServerCert),
		ServerKeyPath:  envOrDefault(lookupEnv, "SENTRY_SERVER_KEY", defaultServerKey),
		CACertPath:     envOrDefault(lookupEnv, "SENTRY_CA_CERT", defaultCACert),
	}

	flags := flag.NewFlagSet("sentry-server", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	flags.IntVar(&config.Port, "port", config.Port, "Port to listen on")
	flags.StringVar(&config.ServerCertPath, "cert", config.ServerCertPath, "Server certificate path")
	flags.StringVar(&config.ServerKeyPath, "key", config.ServerKeyPath, "Server private key path")
	flags.StringVar(&config.CACertPath, "ca", config.CACertPath, "CA certificate path")
	if err := flags.Parse(args); err != nil {
		return serverConfig{}, err
	}
	if flags.NArg() != 0 {
		return serverConfig{}, fmt.Errorf("unexpected arguments: %v", flags.Args())
	}
	if config.Port < 1 || config.Port > 65535 {
		return serverConfig{}, fmt.Errorf("port must be between 1 and 65535, got %d", config.Port)
	}
	for name, path := range map[string]string{
		"server certificate": config.ServerCertPath,
		"server key":         config.ServerKeyPath,
		"CA certificate":     config.CACertPath,
	} {
		if path == "" {
			return serverConfig{}, fmt.Errorf("%s path cannot be empty", name)
		}
	}

	return config, nil
}

func envPort(lookupEnv func(string) (string, bool)) (int, error) {
	value, ok := lookupEnv("SENTRY_PORT")
	if !ok || value == "" {
		return defaultPort, nil
	}

	port, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid SENTRY_PORT %q: %w", value, err)
	}
	return port, nil
}

func envOrDefault(lookupEnv func(string) (string, bool), name, fallback string) string {
	if value, ok := lookupEnv(name); ok && value != "" {
		return value
	}
	return fallback
}
