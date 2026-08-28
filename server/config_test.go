package main

import (
	"reflect"
	"testing"
)

func TestParseServerConfigDefaults(t *testing.T) {
	config, err := parseServerConfig(nil, emptyEnv)
	if err != nil {
		t.Fatalf("parseServerConfig() error = %v", err)
	}

	want := serverConfig{
		Port:           defaultPort,
		ServerCertPath: defaultServerCert,
		ServerKeyPath:  defaultServerKey,
		CACertPath:     defaultCACert,
	}
	if !reflect.DeepEqual(config, want) {
		t.Fatalf("config = %#v, want %#v", config, want)
	}
}

func TestParseServerConfigEnvironmentAndFlagPrecedence(t *testing.T) {
	env := map[string]string{
		"SENTRY_PORT":        "6100",
		"SENTRY_SERVER_CERT": "/env/server.crt",
		"SENTRY_SERVER_KEY":  "/env/server.key",
		"SENTRY_CA_CERT":     "/env/ca.crt",
	}
	lookupEnv := func(name string) (string, bool) {
		value, ok := env[name]
		return value, ok
	}

	config, err := parseServerConfig([]string{
		"-port", "50051",
		"-cert", "/flag/server.crt",
		"-key", "/flag/server.key",
		"-ca", "/flag/ca.crt",
	}, lookupEnv)
	if err != nil {
		t.Fatalf("parseServerConfig() error = %v", err)
	}

	want := serverConfig{
		Port:           50051,
		ServerCertPath: "/flag/server.crt",
		ServerKeyPath:  "/flag/server.key",
		CACertPath:     "/flag/ca.crt",
	}
	if !reflect.DeepEqual(config, want) {
		t.Fatalf("config = %#v, want %#v", config, want)
	}
}

func TestParseServerConfigRejectsInvalidPorts(t *testing.T) {
	for _, port := range []string{"0", "65536", "not-a-port"} {
		t.Run(port, func(t *testing.T) {
			lookupEnv := func(name string) (string, bool) {
				if name == "SENTRY_PORT" {
					return port, true
				}
				return "", false
			}
			if _, err := parseServerConfig(nil, lookupEnv); err == nil {
				t.Fatalf("parseServerConfig() succeeded with SENTRY_PORT=%q", port)
			}
		})
	}
}

func TestParseServerConfigRejectsEmptyTLSPath(t *testing.T) {
	if _, err := parseServerConfig([]string{"-key="}, emptyEnv); err == nil {
		t.Fatal("parseServerConfig() succeeded with an empty key path")
	}
}

func emptyEnv(string) (string, bool) {
	return "", false
}
