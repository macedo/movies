package main

import (
	"testing"

	"github.com/macedo/movies-api/types"
)

func Test_ConnDatabase(t *testing.T) {
	tt := []struct {
		description      string
		env              string
		expectedDatabase string
	}{
		{description: "without env", expectedDatabase: "movies_api_development"},
		{description: "production env", env: "production", expectedDatabase: "movies_api_production"},
		{description: "test env", env: "test", expectedDatabase: "movies_api_test"},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			config := types.Config{Env: tc.env}
			conn, err := NewConnection(config)

			if err != nil {
				t.Fatalf("unexpected error %s", err)
			}

			if conn.Database != tc.expectedDatabase {
				t.Fatalf("expected database %s got %s", tc.expectedDatabase, conn.Database)
			}
		})
	}
}

func Test_PostgresURIForConfigWithURL(t *testing.T) {
	customURL := "my-custom-url"
	config := types.Config{
		Database: types.Database{URL: customURL},
	}

	conn, err := NewConnection(config)

	if err != nil {
		t.Fatalf("unexpected error %s", err)
	}

	output := conn.PostgresURI()
	if output != customURL {
		t.Fatalf("expected URL %s got %s", customURL, output)
	}
}

func Test_PostgresURIForConfigWithoutURL(t *testing.T) {
	config := types.Config{
		Env: "testing",
		Database: types.Database{
			Username: "user",
			Password: "pass",
			Host:     "host",
		},
	}

	conn, err := NewConnection(config)

	if err != nil {
		t.Fatalf("unexpected error %s", err)
	}

	expectedURI := "postgres://user:pass@host/movies_api_testing?sslmode=disable"
	output := conn.PostgresURI()

	if output != expectedURI {
		t.Fatalf("expected URL %s got %s", expectedURI, output)
	}
}
