package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hisbaan/envman/config"
	"github.com/hisbaan/envman/utils"
	"github.com/spf13/cobra"
)

var unlinkCmd = &cobra.Command{
	Use:   "unlink",
	Short: "Remove any .env symlinks",
	RunE: func(cmd *cobra.Command, args []string) error {
		project, err := config.GetCurrentProject()
		if err != nil {
			return err
		}

		paths := utils.Map(project.Dirs, func(dir string) string { return filepath.Join(dir, ".env") })

		// Check that all .env files are not hardlinks
		for _, path := range paths {
			if !utils.IsSymlinkOrDoesNotExist(path) {
				return fmt.Errorf("%s is a hardlink and should not be removed", path)
			}
		}

		for _, path := range paths {
			if _, err := os.Lstat(path); err == nil {
				os.Remove(path)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(unlinkCmd)
}
