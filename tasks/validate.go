package tasks

import "github.com/xeipuuv/gojsonschema"

// ValidationError is returned when the data does not satisfy the JSON schema validations
type ValidationError struct {
	result *gojsonschema.Result
}

// FieldErrors returns a map of invalid fields and their associated lists of errors.
func (ve *ValidationError) FieldErrors() map[string][]string {
	fe := map[string][]string{}

	for _, e := range ve.result.Errors() {
		fieldName := e.Field()
		errors, ok := fe[fieldName]
		if !ok {
			errors = []string{}
		}
		errors = append(errors, e.Description())
		fe[fieldName] = errors
	}

	return fe
}

// Error implements the error interface for ValidationError
func (ve *ValidationError) Error() string {
	return "the data is not valid"
}

// Validate checks the data for a task against it's JSON schema
func (t *Task) Validate(data map[string]interface{}) error {
	schemaLoader := gojsonschema.NewGoLoader(t.Schema)
	dataLoader := gojsonschema.NewGoLoader(data)

	result, err := gojsonschema.Validate(schemaLoader, dataLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {
		return &ValidationError{result: result}
	}

	return nil
}
