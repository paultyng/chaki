package server

import (
	"crypto/rsa"
	"fmt"

	"github.com/labstack/echo"
	glog "github.com/labstack/gommon/log"
	"golang.org/x/oauth2"

	"github.com/paultyng/chaki/server/sso"
	"github.com/paultyng/chaki/tasks"
)

const githubEnterpriseDomain = "github.ua.com"

// Config represents the config for a new Server.
type Config struct {
	Tasks             *tasks.Config
	PrivateKey        *rsa.PrivateKey
	OAuthClientID     string
	OAuthClientSecret string
}

// Server represents the state of the web server.
type Server struct {
	echo        *echo.Echo
	tasksConfig *tasks.Config
}

// New creates a new instance of the web server.
func New(c *Config) *Server {
	s := &Server{
		echo:        echo.New(),
		tasksConfig: c.Tasks,
	}

	e := s.echo
	e.Logger.SetLevel(glog.INFO)
	//e.Use(middleware.Logger())

	ssoConf := &sso.OAuth2Config{
		JWTAuthConfig: sso.JWTAuthConfig{
			PrivateKey: c.PrivateKey,
		},
		OAuth2: &oauth2.Config{
			ClientID:     c.OAuthClientID,
			ClientSecret: c.OAuthClientSecret,
			Scopes:       []string{"user:email"},
			Endpoint:     sso.GithubEnterpriseEndpoint(githubEnterpriseDomain),
		},
		EmailLookupFunc: sso.GithubEnterpriseEmailLookup(githubEnterpriseDomain),
		NoAuthn:         true,
	}

	e.Use(sso.OAuth2FromConfig(ssoConf))

	api := e.Group("/api")
	api.Use(sso.JWTAuthFromConfig(&ssoConf.JWTAuthConfig))
	tasks := api.Group("/tasks")
	tasks.GET("", s.getTasks)
	tasks.POST("/:name/run", s.runTask)

	names := s.tasksConfig.TaskNames()
	for _, n := range names {
		e.GET(fmt.Sprintf("/%s", n), serveStatic("build/index.html"))
	}

	e.GET("/*", serveStatic("build"))

	return s
}

// Start makes the server start listening on the specified binding.
func (s *Server) Start(binding string) error {
	return s.echo.Start(binding)
}
