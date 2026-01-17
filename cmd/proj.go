package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hisbaan/envman/config"
	"github.com/hisbaan/envman/utils"
	"github.com/spf13/cobra"
)

func getProjectPath(args []string) (string, error) {
	if len(args) < 1 {
		return os.Getwd()
	}
	return filepath.Abs(args[0])
}

var projAddCmd = &cobra.Command{
	Use:   "add [project-dir]",
	Short: "Add the specified dir as a project, falling back to $CWD",
	RunE: func(cmd *cobra.Command, args []string) error {
		project, err := getProjectPath(args)
		if err != nil {
			return err
		}

		_, err = config.GetProject(project)
		if err == nil {
			return fmt.Errorf("project '%s' already exists", project)
		}

		fmt.Printf("added project %q with root directory\n", project)
		if err := config.AddProject(project); err != nil {
			return err
		}

		return nil
	},
}

var projRmCmd = &cobra.Command{
	Use:   "rm [project-dir]",
	Short: "Remove the specified dir as a project, falling back to $CWD",
	RunE: func(cmd *cobra.Command, args []string) error {
		project, err := getProjectPath(args)
		if err != nil {
			return err
		}

		_, err = config.GetProject(project)
		if err != nil {
			return fmt.Errorf("project '%s' does not exist", project)
		}

		yes, _ := cmd.Flags().GetBool("yes")
		if !yes && !utils.Confirm(fmt.Sprintf("Remove project %q?", project), false) {
			return nil
		}

		if err := config.RemoveProject(project); err != nil {
			return err
		}

		return nil
	},
}

var projCmd = &cobra.Command{
	Use:   "proj",
	Short: "Add or remove the specified project",
}

// TODO completions for add/rm and [project]
func init() {
	projRmCmd.Flags().BoolP("yes", "y", false, "skip confirmation")

	projCmd.AddCommand(projAddCmd)
	projCmd.AddCommand(projRmCmd)

	projCmd.Aliases = append(projCmd.Aliases, "project")

	rootCmd.AddCommand(projCmd)
}
