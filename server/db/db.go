package db

import (
	"database/sql"
	"net/url"

	_ "github.com/lib/pq"
)

type Connection struct {
	DbName         string
	User, Password string
	Host           string
	DisableSSL     bool
}

func (c *Connection) ConnectionURL() string {
	dbUrl := &url.URL{
		Scheme: "postgres", // this field is actually driver name and should always be "postgres"
		Host:   c.Host,
		User:   url.UserPassword(c.User, c.Password),
		Path:   c.DbName,
	}
	if c.DisableSSL {
		dbUrl.RawQuery = url.Values{
			"sslmode":     []string{"disable"},
			"search_path": []string{"plants"},
		}.Encode()
	} else {
		dbUrl.RawQuery = url.Values{
			"search_path": []string{"plants"},
		}.Encode()
	}
	return dbUrl.String()
}

func (c *Connection) Open() (*sql.DB, error) {
	return sql.Open("postgres", c.ConnectionURL())
}
