package tasks

// Sanitize returns a copy of the config sanitized for client side input.
// This is strictly a whitelist for safety.
func (src *Config) Sanitize() *Config {
	dst := &Config{
		Tasks: map[string]Task{},
	}

	for name, srcT := range src.Tasks {
		dstT := Task{
			Title:       srcT.Title,
			Description: srcT.Description,
			Schema:      srcT.Schema,
			UISchema:    srcT.UISchema,
		}

		dst.Tasks[name] = dstT
	}

	return dst
}
