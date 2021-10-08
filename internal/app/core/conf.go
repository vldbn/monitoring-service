package core

import (
	"fmt"
	"os"
)

// Config application config struct
type Config struct {
	username    string
	password    string
	secretKey   string
	port        string
	database    string
	databaseURL string
}

// Username getter
func (c *Config) Username() string {
	return c.username
}

// SetUsername setter
func (c *Config) SetUsername(username string) {
	c.username = username
}

// Password getter
func (c *Config) Password() string {
	return c.password
}

// SecretKey getter
func (c *Config) SecretKey() string {
	return c.secretKey
}

// Port getter
func (c *Config) Port() string {
	return c.port
}

// Database getter
func (c *Config) Database() string {
	return c.database
}

// DatabaseURL getter
func (c *Config) DatabaseURL() string {
	return c.databaseURL
}

// NewConfig constructor
func NewConfig() *Config {
	p := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if p == ":" {
		p = ":8000"
	}
	sk := os.Getenv("SECRET_KEY")
	if sk == "" {
		sk = "secretKeyString"
	}
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "mongodb://localhost:27017"
	}
	usr := os.Getenv("USERNAME")
	if usr == "" {
		usr = "username"
	}
	pwd := os.Getenv("PASSWORD")
	if pwd == "" {
		pwd = "password"
	}
	return &Config{
		username:    usr,
		password:    pwd,
		secretKey:   sk,
		port:        p,
		database:    "monitoring",
		databaseURL: dbUrl,
	}
}
