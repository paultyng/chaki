package tasks

import (
	"io"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// Config represents the task configuration file
type Config struct {
	Tasks         map[string]Task         `json:"tasks"`
	DBConnections map[string]DBConnection `json:"dbConnections"`
}

// DBTask represents the database specific parameters for a task
type DBTask struct {
	Connection string `json:"connection"`
	SQL        string `json:"sql"`
}

// DBConnection represents the connection info for a DB task
type DBConnection struct {
	Driver     string `json:"driver"`
	DataSource string `json:"dataSource"`
}

// Task represents a task to execute
type Task struct {
	Name        string                 `json:"name"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Schema      map[string]interface{} `json:"schema"`
	UISchema    map[string]interface{} `json:"uiSchema"`
	DB          *DBTask                `json:"db"`
}

// NewConfig creates a Config by unmarshaling YAML or JSON from the Reader
func NewConfig(r io.Reader) (*Config, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal(bytes, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
