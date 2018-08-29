package main

import (
	"fmt"
	"os/user"

	"github.com/macedo/movies-api/types"
)

type Conn struct {
	Database string

	username string
	password string
	host     string
	url      string
}

const (
	defaultENV    = "development"
	defaultPrefix = "movies_api"
)

func NewConnection(c types.Config) (Conn, error) {
	conn := Conn{
		password: c.Database.Password,
		host:     c.Database.Host,
		url:      c.Database.URL,
	}

	env := c.Env
	if env == "" {
		env = defaultENV
	}

	username := c.Database.Username
	if username == "" {
		currentUser, err := user.Current()

		if err != nil {
			return Conn{}, err
		}

		username = currentUser.Username
	}

	conn.username = username
	conn.Database = fmt.Sprintf("%s_%s", defaultPrefix, env)

	return conn, nil
}

// PostgresURI build a postgres URI based on conn information
func (c Conn) PostgresURI() string {
	if c.url != "" {
		return c.url
	}

	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.username, c.password, c.host, c.Database)
}
