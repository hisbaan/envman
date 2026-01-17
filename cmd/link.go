package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/hisbaan/envman/config"
	"github.com/hisbaan/envman/utils"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link [environment]",
	Short: "Link the provided environment to .env",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("environment must be specified")
		}
		env := args[0]

		project, err := config.GetCurrentProject()
		if err != nil {
			return err
		}

		envs, err := utils.GetCommonEnvs(project.Dirs)
		if err != nil {
			return err
		}

		if !slices.Contains(envs, env) {
			return fmt.Errorf("provided environment must be one of %s", strings.Join(envs, ", "))
		}

		paths := utils.Map(project.Dirs, func(dir string) string { return filepath.Join(dir, ".env") })

		// Check that all .env files are not hardlinks
		for _, path := range paths {
			if !utils.IsSymlinkOrDoesNotExist(path) {
				return fmt.Errorf("%s is a hardlink and should not be overwritten", path)
			}
		}

		for _, path := range paths {
			if _, err := os.Lstat(path); err == nil {
				os.Remove(path)
			}
			err := os.Symlink(path+"."+env, path)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	// TODO define some custom completions that fetch the values returned by utils.GetCommonEnvs
	rootCmd.AddCommand(linkCmd)
}
