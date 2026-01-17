package cmd

import (
	"fmt"
	"path/filepath"
	"slices"

	"github.com/hisbaan/envman/config"
	"github.com/hisbaan/envman/utils"
	"github.com/spf13/cobra"
)

var dirAddCmd = &cobra.Command{
	Use:   "add [dir]",
	Short: "Adds the specified dir to the associated project",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("dir must be specified")
		}
		dir, err := filepath.Abs(args[0])
		if err != nil {
			return err
		}

		project, err := config.GetParentProject(dir)
		if err != nil {
			return fmt.Errorf("directory is not a child of any defined projects")
		}

		if slices.Contains(project.Dirs, dir) {
			return fmt.Errorf("directory already defined in project %q", project.Path)
		}

		err = config.AddDir(project.Path, dir)
		if err != nil {
			return err
		}
		fmt.Printf("Added directory %q to project %q\n", dir, project.Path)

		return nil
	},
}

var dirRmCmd = &cobra.Command{
	Use:   "rm [dir]",
	Short: "Removes the specified dir to the associated project",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("dir must be specified")
		}
		dir, err := filepath.Abs(args[0])
		if err != nil {
			return err
		}

		project, err := config.GetParentProject(dir)
		if err != nil {
			return fmt.Errorf("directory is not a child of any defined projects")
		}

		if !slices.Contains(project.Dirs, dir) {
			return fmt.Errorf("directory is not defined in project %q", project.Path)
		}

		yes, _ := cmd.Flags().GetBool("yes")
		if !yes && !utils.Confirm(fmt.Sprintf("Remove directory %q?", dir), false) {
			return nil
		}

		err = config.RemoveDir(project.Path, dir)
		if err != nil {
			return err
		}

		return nil
	},
}

var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Add or remove the specified directory from the current project",
}

// TODO completions for add/rm and [dir]
func init() {
	dirRmCmd.Flags().BoolP("yes", "y", false, "skip confirmation")

	dirCmd.AddCommand(dirAddCmd)
	dirCmd.AddCommand(dirRmCmd)

	rootCmd.AddCommand(dirCmd)
}
