package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"

	"github.com/paultyng/chaki/tasks"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:        "execute",
	Aliases:    []string{"exec", "e"},
	SuggestFor: []string{"run"},
	Short:      "Execute a task from the command line",
	Long: `Execute allows you to run a task without the need
for the web UI.  Simply supply data via JSON or YAML, as a
parameter or on STDIN, and the CLI will handle the rest.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		taskName, err := cmd.PersistentFlags().GetString("name")
		if err != nil {
			return err
		}
		if len(taskName) == 0 {
			return fmt.Errorf("a task name is required")
		}

		data, err := cmd.PersistentFlags().GetString("data")
		if err != nil {
			return err
		}

		var r io.Reader

		if data == "-" {
			log.Printf("[INFO] Using data from stdin")
			r = os.Stdin
		} else if len(data) > 0 {
			r = strings.NewReader(data)
		}

		var taskData map[string]interface{}

		if r != nil {
			yamlData, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}

			err = yaml.Unmarshal(yamlData, &taskData)
			if err != nil {
				return err
			}
		}

		config, err := loadTaskConfig(cmd)

		log.Printf("[INFO] Running task %s", taskName)

		res, err := config.Run(taskName, taskData)
		if ve, ok := err.(*tasks.ValidationError); ok {
			for field, errors := range ve.FieldErrors() {
				for _, e := range errors {
					log.Printf("[ERROR] %s: %s", field, e)
				}
			}
			return ve
		}
		if err != nil {
			return err
		}

		log.Printf("[INFO] Execute success!")

		switch t := res.(type) {
		case *tasks.DBTaskResult:
			//TODO: output data
		default:
			log.Printf("[WARN] Unexected result type %T", t)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(execCmd)

	execCmd.PersistentFlags().StringP("name", "n", "", "name of the task to execute")
	execCmd.PersistentFlags().StringP("data", "d", "", "JSON data for the task, use '-' for stdin")

	addTaskConfigFlags(execCmd)
}
