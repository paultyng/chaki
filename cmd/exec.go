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

	"go.ua-ecm.com/chaki/tasks"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		tasksFile, err := cmd.PersistentFlags().GetString("tasks-file")
		if err != nil {
			return err
		}

		log.Printf("[INFO] Using task file %s", tasksFile)

		file, err := os.Open(tasksFile)
		if err != nil {
			return err
		}

		config, err := tasks.NewConfig(file)
		if err != nil {
			return err
		}

		log.Printf("[INFO] Running task %s", taskName)

		err = config.Run(taskName, taskData)
		if ve, ok := err.(*tasks.ValidationError); ok {
			for _, e := range ve.Result.Errors() {
				log.Printf("[Error] %s", e)
			}
			return ve
		}
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(execCmd)

	execCmd.PersistentFlags().StringP("name", "n", "", "name of the task to execute")
	execCmd.PersistentFlags().StringP("data", "d", "", "JSON data for the task, use '-' for stdin")

	execCmd.PersistentFlags().StringP("tasks-file", "t", "tasks.yaml", "file for loading task list")
}
