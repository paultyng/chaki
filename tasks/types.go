package tasks

import (
	"io"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// OptionalStringArray represents an array that can be serialized either as a
// single string or an array of strings
type OptionalStringArray []string

// Config represents the task configuration file
type Config struct {
	Tasks         map[string]Task         `json:"tasks"`
	DBConnections map[string]DBConnection `json:"dbConnections,omitempty"`
}

// DBTask represents the database specific parameters for a task
type DBTask struct {
	Connection string              `json:"connection"`
	SQL        OptionalStringArray `json:"sql"`
}

// DBConnection represents the connection info for a DB task
type DBConnection struct {
	Driver     string `json:"driver"`
	DataSource string `json:"dataSource"`
}

// Task represents a task to execute
type Task struct {
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Schema      map[string]interface{} `json:"schema"`
	UISchema    map[string]interface{} `json:"uiSchema"`
	DB          *DBTask                `json:"db,omitempty"`
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

// Task returns a config task by name if it exists, otherwise nil.
func (c *Config) Task(name string) *Task {
	t, ok := c.Tasks[name]
	if !ok {
		return nil
	}

	return &t
}

// TaskNames returns the names of all tasks registered in the Config.
func (c *Config) TaskNames() []string {
	if c.Tasks == nil {
		return []string{}
	}

	names := make([]string, len(c.Tasks))

	i := 0
	for name := range c.Tasks {
		names[i] = name
		i++
	}

	return names
}
