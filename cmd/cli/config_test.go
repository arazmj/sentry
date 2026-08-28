package main

import (
	"reflect"
	"testing"
)

func TestParseCLIConfigDefaults(t *testing.T) {
	config, args, err := parseCLIConfig([]string{"list"}, emptyEnv)
	if err != nil {
		t.Fatalf("parseCLIConfig() error = %v", err)
	}

	want := cliConfig{
		ServerAddr:     defaultServerAddr,
		ClientCertPath: defaultClientCert,
		ClientKeyPath:  defaultClientKey,
		CACertPath:     defaultCACert,
	}
	if !reflect.DeepEqual(config, want) {
		t.Fatalf("config = %#v, want %#v", config, want)
	}
	if !reflect.DeepEqual(args, []string{"list"}) {
		t.Fatalf("args = %v, want [list]", args)
	}
}

func TestParseCLIConfigEnvironmentAndFlagPrecedence(t *testing.T) {
	env := map[string]string{
		"SENTRY_SERVER":      "env.example:6000",
		"SENTRY_CLIENT_CERT": "/env/client.crt",
		"SENTRY_CLIENT_KEY":  "/env/client.key",
		"SENTRY_CA_CERT":     "/env/ca.crt",
	}
	lookupEnv := func(name string) (string, bool) {
		value, ok := env[name]
		return value, ok
	}

	config, args, err := parseCLIConfig([]string{
		"-server", defaultServerAddr,
		"-cert", "/flag/client.crt",
		"-key", "/flag/client.key",
		"-ca", "/flag/ca.crt",
		"start", "-cmd", "echo",
	}, lookupEnv)
	if err != nil {
		t.Fatalf("parseCLIConfig() error = %v", err)
	}

	want := cliConfig{
		ServerAddr:     defaultServerAddr,
		ClientCertPath: "/flag/client.crt",
		ClientKeyPath:  "/flag/client.key",
		CACertPath:     "/flag/ca.crt",
	}
	if !reflect.DeepEqual(config, want) {
		t.Fatalf("config = %#v, want %#v", config, want)
	}
	if !reflect.DeepEqual(args, []string{"start", "-cmd", "echo"}) {
		t.Fatalf("args = %v, want [start -cmd echo]", args)
	}
}

func TestParseCLIConfigRejectsEmptyTLSPath(t *testing.T) {
	if _, _, err := parseCLIConfig([]string{"-cert=", "list"}, emptyEnv); err == nil {
		t.Fatal("parseCLIConfig() succeeded with an empty certificate path")
	}
}

func emptyEnv(string) (string, bool) {
	return "", false
}
