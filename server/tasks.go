package server

import (
	"net/http"

	"go.ua-ecm.com/chaki/tasks"

	"github.com/labstack/echo"
)

func (s *Server) getTasks(c echo.Context) error {
	sanitizedConfig := s.tasksConfig.Sanitize()

	return c.JSON(http.StatusOK, sanitizedConfig)
}

type runTaskResponseStatement struct {
	Data []map[string]interface{} `json:"data,omitempty"`
}

type runTaskResponse struct {
	Statements []runTaskResponseStatement `json:"statements,omitempty"`
}

func (s *Server) runTask(c echo.Context) error {
	log := c.Logger()

	req := &struct {
		Data map[string]interface{} `json:"data"`
	}{}

	if err := c.Bind(req); err != nil {
		return err
	}

	name := c.Param("name")
	log.Infof("running task %s", name)
	result, err := s.tasksConfig.Run(name, req.Data)
	if err != nil {
		return err
	}

	var resp *runTaskResponse

	switch t := result.(type) {
	case *tasks.DBTaskResult:
		dbres := result.(*tasks.DBTaskResult)
		resp = &runTaskResponse{
			Statements: make([]runTaskResponseStatement, len(dbres.Statements)),
		}
		for i, sr := range dbres.Statements {
			resp.Statements[i].Data = sr.Data
		}

	default:
		log.Warnf("unexpected result type %T", t)
	}

	if resp == nil {
		resp = &runTaskResponse{}
	}

	return c.JSON(http.StatusOK, resp)
}
