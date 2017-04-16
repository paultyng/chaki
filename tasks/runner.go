package tasks

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

// ValidationError is returned when the data does not satisfy the JSON schema validations
type ValidationError struct {
	Result *gojsonschema.Result
}

// Error implements the error interface for ValidationError
func (ve *ValidationError) Error() string {
	return "the data is not valid"
}

// Run executes a task by name from a given config using the specified data
func (c *Config) Run(name string, data map[string]interface{}) error {
	t, ok := c.Tasks[name]
	if !ok {
		return fmt.Errorf("unable to find task %s", name)
	}

	schemaLoader := gojsonschema.NewGoLoader(t.Schema)
	dataLoader := gojsonschema.NewGoLoader(data)

	result, err := gojsonschema.Validate(schemaLoader, dataLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {
		return &ValidationError{Result: result}
	}

	return t.run(data, c)
}

func (t *Task) run(data map[string]interface{}, c *Config) error {
	switch {
	case t.DB != nil:
		return t.DB.run(data, c)
	default:
		return fmt.Errorf("no supported task subtypes")
	}
}
