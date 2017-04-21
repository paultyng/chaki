package tasks

import "fmt"

// Run executes a task by name from a given config using the specified data
func (c *Config) Run(name string, data map[string]interface{}) (interface{}, error) {
	t := c.Task(name)
	if t == nil {
		return nil, fmt.Errorf("unable to find task %s", name)
	}

	err := t.Validate(data)
	if err != nil {
		return nil, err
	}

	return t.run(data, c)
}

func (t *Task) run(data map[string]interface{}, c *Config) (interface{}, error) {
	switch {
	case t.DB != nil:
		return t.DB.run(data, c)
	default:
		return nil, fmt.Errorf("no supported task subtypes")
	}
}
