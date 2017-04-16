package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func (s *Server) getTasks(c echo.Context) error {
	sanitizedConfig := s.tasksConfig.Sanitize()

	return c.JSON(http.StatusOK, sanitizedConfig)
}

type runTaskRequest struct {
	Data map[string]interface{} `json:"data"`
}

func (s *Server) runTask(c echo.Context) error {
	log := c.Logger()

	req := &runTaskRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}

	name := c.Param("name")
	log.Infof("running task %s", name)
	return s.tasksConfig.Run(name, req.Data)
}
