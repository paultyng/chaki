package server

import (
	"fmt"

	"github.com/labstack/echo"
	glog "github.com/labstack/gommon/log"
	"go.ua-ecm.com/chaki/tasks"
)

// Server represents the state of the web server.
type Server struct {
	echo        *echo.Echo
	tasksConfig *tasks.Config
}

// New creates a new instance of the web server.
func New(c *tasks.Config) *Server {
	s := &Server{
		echo:        echo.New(),
		tasksConfig: c,
	}
	e := s.echo
	e.Logger.SetLevel(glog.INFO)

	e.GET("/api/tasks", s.getTasks)
	e.POST("/api/tasks/:name/run", s.runTask)

	names := s.tasksConfig.TaskNames()
	for _, n := range names {
		e.Static(fmt.Sprintf("/%s", n), "build")
	}

	e.Static("/", "build")

	return s
}

// Start makes the server start listening on the specified binding.
func (s *Server) Start(binding string) error {
	return s.echo.Start(binding)
}
